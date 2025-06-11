package blocks

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type LastError struct {
	Message string `json:"message"`
}

type Block struct {
	ID                uuid.UUID       `json:"id"`
	Type              DataType        `json:"type"`
	OriginalSpaceID   uuid.UUID       `json:"original_space_id"`
	RootParentID      *uuid.UUID      `json:"root_parent_id"`
	ParentID          *uuid.UUID      `json:"parent_id,omitempty"`
	AccountID         uuid.UUID       `json:"account_id"`
	SpaceID           uuid.UUID       `json:"space_id"`
	PreviousSpaceID   uuid.UUID       `json:"previous_space_id"`
	CreatorUserID     uuid.UUID       `json:"creator_user_id"`
	Properties        Properties      `json:"properties"`
	Styles            Properties      `json:"styles"`
	Content           []uuid.UUID     `json:"content"`
	ChildrenRecursive []uuid.UUID     `json:"children_recursive"`
	RawBody           string          `json:"raw_body"` // html, email, etc
	LifecycleStatus   LifecycleStatus `json:"lifecycle_status"`
	Metadata          json.RawMessage `json:"metadata,omitempty"`
	Origin            Origin          `json:"origin"`
	MovesHistory      []Move          `json:"moves_history"`
	LastError         LastError       `json:"last_error"`
	LastViewedAt      *time.Time      `json:"last_viewed_at,omitempty"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`

	// Meaning contains a short explanation of what this block is about.
	// This field should be set only for root blocks (blocks without parents).
	Meaning string `json:"meaning,omitempty"`

	// Classification contains the classification data for the block
	Classification Classification `json:"classification,omitempty"`

	// generated content for embeddings
	CalculatedContent string

	// For search purposes
	DenseVector []float32 `json:"dense_vector"`
}

func (b *Block) IsEmpty() bool {
	// Check if there's any content
	if len(b.Content) > 0 {
		return false
	}

	// Check if there are any properties
	if len(b.Properties) == 0 {
		return true
	}

	// Check if all property values are empty
	for _, values := range b.Properties {
		if len(values) == 0 {
			continue
		}

		for _, value := range values {
			if value == nil {
				continue
			}

			switch v := value.(type) {
			case string:
				if v != "" {
					return false
				}
			default:
				// For any other type, consider it non-empty
				return false
			}
		}
	}

	return true
}

func (b *Block) LockKeyForEnrichment() string {
	return fmt.Sprintf("block_enrichment_%s", b.ID.String())
}

func (b *Block) DebounceKeyForEditing() string {
	return DebounceKeyForEditing(b.ID)
}

func DebounceKeyForEditing(blockID uuid.UUID) string {
	return fmt.Sprintf("block_editing_debounce_%s", blockID.String())
}

func TagParent(blockID uuid.UUID) string {
	return fmt.Sprintf("parent_block_%s", blockID.String())
}

func (b *Block) AnnotationID() string {
	return fmt.Sprintf("[[block:%s]]", b.ID.String())
}

// ExtractReferences extracts all UUIDs from a string that match the [[block:uuid]] pattern
func ExtractReferences(text string) []uuid.UUID {
	// Regex to match [[block:uuid]] pattern
	pattern := regexp.MustCompile(`\[\[block:([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})\]\]`)

	// Find all matches
	matches := pattern.FindAllStringSubmatch(text, -1)

	// Extract UUIDs from matches
	uuids := make([]uuid.UUID, 0, len(matches))
	for _, match := range matches {
		if len(match) >= 2 {
			// Parse UUID from the captured group
			id, err := uuid.Parse(match[1])
			if err == nil {
				uuids = append(uuids, id)
			}
		}
	}

	return uuids
}

func (b *Block) Keywords() []string {
	uniqueKeywords := make(map[string]bool)
	for _, move := range b.MovesHistory {
		for _, keyword := range move.SpaceKeywords {
			uniqueKeywords[keyword] = true
		}
		for _, keyword := range move.ReasoningKeywords {
			uniqueKeywords[keyword] = true
		}
	}

	var keywords []string
	for keyword := range uniqueKeywords {
		keywords = append(keywords, keyword)
	}
	return keywords
}

func (b *Block) AppendChild(id uuid.UUID) {
	b.Content = append(b.Content, id)
}

func (b *Block) RemoveChild(id uuid.UUID) error {
	for i, childID := range b.Content {
		if childID == id {
			// Remove the child by slicing it out of the Content array
			b.Content = append(b.Content[:i], b.Content[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("child with ID %s not found", id.String())
}

func (b *Block) InsertChild(id uuid.UUID, afterID uuid.UUID) error {
	// Find the position of afterID in the content array
	position := -1
	for i, childID := range b.Content {
		if childID == afterID {
			position = i
			break
		}
	}

	if position == -1 {
		return fmt.Errorf("reference element with ID %s not found", afterID.String())
	}

	// Insert the new child after the specified position
	if position == len(b.Content)-1 {
		// If afterID is the last element, simply append
		b.Content = append(b.Content, id)
	} else {
		// Insert the new element after the specified position
		b.Content = append(b.Content[:position+1], append([]uuid.UUID{id}, b.Content[position+1:]...)...)
	}

	return nil
}

func (b *Block) CheckPermissions(accountID uuid.UUID) error {
	if b.AccountID != accountID {
		return ErrUnauthorizedBlockAccess
	}
	return nil
}

func (b *Block) GetFirstNContent(n int) []uuid.UUID {
	if len(b.Content) >= n {
		return b.Content[0:n]
	}
	return b.Content
}

// AddMove adds a move record to the block's move history
// It maintains a maximum of 20 most recent moves
func (b *Block) AddMove(move Move) {
	// Keep only last 20 moves
	if len(b.MovesHistory) >= 20 {
		// Remove first element by slicing from index 1 onwards
		b.MovesHistory = b.MovesHistory[1:]
	}
	b.MovesHistory = append(b.MovesHistory, move)
}

// UpdateFromJSON updates the Block from JSON data and returns a list of updated field names.
// It only updates fields that exist in the JSON and differ from the current Block.
func (b *Block) UpdateFromJSON(data []byte) ([]string, error) {
	// First unmarshal into a map to get the raw JSON properties
	var rawData map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawData); err != nil {
		return nil, err
	}

	// Track which fields were updated
	updatedFields := []string{}

	// Create a map of JSON field names to struct fields
	jsonToStructField := make(map[string]string)

	// Use reflection to build the map and handle the update
	v := reflect.ValueOf(b).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonFieldName := getJSONFieldName(field)
		jsonToStructField[jsonFieldName] = field.Name
	}

	// Create a temporary Block to store the current state for comparison
	currentBlock := *b

	// Iterate through the raw JSON data
	for jsonField, rawValue := range rawData {
		// Skip UpdatedAt as we'll set it at the end
		if jsonField == "updated_at" {
			continue
		}

		// Look up the corresponding struct field
		structFieldName, exists := jsonToStructField[jsonField]
		if !exists {
			// JSON field doesn't correspond to any struct field, skip it
			continue
		}

		// Get the field by name
		fieldIndex := -1
		for i := 0; i < t.NumField(); i++ {
			if t.Field(i).Name == structFieldName {
				fieldIndex = i
				break
			}
		}

		if fieldIndex == -1 {
			// Field not found in struct, skip it
			continue
		}

		// Get the field and its type
		field := v.Field(fieldIndex)
		fieldType := field.Type()

		// Create a new instance of the field's type to unmarshal into
		newValue := reflect.New(fieldType).Interface()

		// Unmarshal the raw JSON value into the new instance
		if err := json.Unmarshal(rawValue, newValue); err != nil {
			// Skip fields that can't be unmarshaled
			continue
		}

		// Extract the value from the pointer
		newValueElem := reflect.ValueOf(newValue).Elem()

		// Skip if the field has zero value
		if reflect.DeepEqual(newValueElem.Interface(), reflect.Zero(fieldType).Interface()) {
			continue
		}

		// Get current value for comparison
		currentField := reflect.ValueOf(currentBlock).FieldByName(structFieldName)

		// Only update and track fields that have changed
		if !reflect.DeepEqual(currentField.Interface(), newValueElem.Interface()) {
			field.Set(newValueElem)
			updatedFields = append(updatedFields, jsonField)
		}
	}

	// Update the UpdatedAt field
	if len(updatedFields) > 0 {
		b.UpdatedAt = time.Now()
		updatedFields = append(updatedFields, "updated_at")
	}

	return updatedFields, nil
}

func NewEmptyBlock() Block {
	return Block{
		ID:         uuid.New(),
		Type:       TypeFragment,
		CreatedAt:  time.Now(),
		Properties: Properties{},
		Styles:     Properties{},
	}
}

// getJSONFieldName extracts the JSON field name from struct field tags
func getJSONFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if tag == "" {
		return field.Name
	}

	// Parse the json tag which might have options like "name,omitempty"
	parts := strings.Split(tag, ",")
	if parts[0] == "-" {
		// Field is explicitly excluded from JSON
		return field.Name
	}
	if parts[0] != "" {
		return parts[0]
	}

	return field.Name
}

func RenderProperties(ctx context.Context, b Block) string {
	switch b.Type {
	case TypeParagraph:
		return RenderParagraphProperties(b)
	case TypePage:
		// For page and text blocks, also use the title property if available
		title, ok := b.Properties.Get(PropertyKeyTitle)
		if !ok {
			return ""
		}
		titleString, ok := title.(string)
		if !ok {
			return ""
		}
		return titleString
	case TypeLink:
		return RenderLinkProperties(b)
	case TypeToDo:
		return RenderToDoProperties(b)
	case TypeYouTube:
		return RenderYoutubeProperties(b)
	case TypeMovie:
		return RenderMovieProperties(b)
	case TypeSeries:
		return RenderSeriesProperties(b)
	case TypeEmail:
		return RenderEmailProperties(b)
	case TypeInstagram:
		return RenderInstagramProperties(b)
	case TypeBook:
		return RenderBookProperties(b)
	case TypeLine:
		return RenderLineProperties(b)
	case TypePerson:
		return RenderPersonProperties(b)
	case TypePlace:
		return RenderPlaceProperties(b)
	case TypeHeader1, TypeHeader2, TypeHeader3, TypeHeader4, TypeHeader5, TypeHeader6:
		return RenderHeaderProperties(b)
	case TypeNumberedListItem:
		return RenderNumberedListItemProperties(b)
	case TypeBulletListItem:
		return RenderBulletListItemProperties(b)
	default:
		return ""
	}
}
