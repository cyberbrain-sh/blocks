package blocks

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// main object properties
const (
	BlockPropertyID                = "id"
	BlockPropertyType              = "type"
	BlockPropertyRootParentID      = "root_parent_id"
	BlockPropertyOriginalSpaceID   = "original_space_id"
	BlockPropertyParentID          = "parent_id"
	BlockPropertyAccountID         = "account_id"
	BlockPropertySpaceID           = "space_id"
	BlockPropertyPreviousSpaceID   = "previous_space_id"
	BlockPropertyCreatorUserID     = "creator_user_id"
	BlockPropertyProperties        = "properties"
	BlockPropertyStypes            = "styles"
	BlockPropertyContent           = "content"
	BlockPropertyMeaning           = "meaning"
	BlockPropertyClassification    = "classification"
	BlockPropertyChildrenRecursive = "children_recursive"
	BlockPropertyRawBody           = "raw_body"
	BlockPropertyLifecycleStatus   = "lifecycle_status"
	BlockPropertyMetadata          = "metadata"
	BlockPropertyOrigin            = "origin"
	BlockPropertyMovesHistory      = "moves_history"
	BlockPropertyLastError         = "last_error"
	BlockPropertyLastViewedAt      = "last_viewed_at"
	BlockPropertyCreatedAt         = "created_at"
	BlockPropertyUpdatedAt         = "updated_at"
)

// PropertyType defines the expected type for a property value
type PropertyType int

const (
	TypeString PropertyType = iota
	TypeInt
	TypeFloat
	TypeBool
	TypeDateTime
	TypeStringArray
	TypeFloatArray // For coordinates array [latitude, longitude]
	TypeAny        // For backward compatibility
)

// Common properties
const PropertyKeyTitle string = "title"
const PropertyKeyText string = "text"
const PropertyKeyDescription string = "description"
const PropertyKeyTargetDateTime string = "target_datetime"
const PropertyKeyReminderOffset string = "reminder_offset"
const PropertyKeyChecked string = "checked"
const PropertyKeyEnriched string = "enriched"

// URL/Link properties
const PropertyKeyURL string = "url"
const PropertyKeyImageURL string = "url_image"

// Movie/TV Series properties
const PropertyKeyIMDBID string = "imdb_id"
const PropertyKeyTMDBID string = "tmdb_id"
const PropertyKeyReleaseYear string = "release_year"
const PropertyKeyFirstAirYear string = "first_air_year"
const PropertyKeyLastAirYear string = "last_air_year"
const PropertyKeyRating string = "rating"
const PropertyKeyRuntime string = "runtime"
const PropertyKeyGenres string = "genres"
const PropertyKeyDirectors string = "directors"
const PropertyKeyCreators string = "creators"
const PropertyKeyCast string = "cast"
const PropertyKeyTagline string = "tagline"
const PropertyKeyBudget string = "budget"
const PropertyKeyRevenue string = "revenue"
const PropertyKeyStatus string = "status"
const PropertyKeyInProduction string = "in_production"
const PropertyKeyNumberOfSeasons string = "number_of_seasons"
const PropertyKeyNumberOfEpisodes string = "number_of_episodes"
const PropertyKeyNetworks string = "networks"
const PropertyKeyType string = "type"

// YouTube properties
const PropertyKeyVideoID string = "video_id"
const PropertyKeyChannelID string = "channel_id"
const PropertyKeyChannelTitle string = "channel_title"
const PropertyKeyPublishedAt string = "published_at"
const PropertyKeyViewCount string = "view_count"
const PropertyKeyLikeCount string = "like_count"
const PropertyKeyCommentCount string = "comment_count"
const PropertyKeyDuration string = "duration"
const PropertyKeyDefinition string = "definition"
const PropertyKeyHasCaptions string = "has_captions"
const PropertyKeyTags string = "tags"

// Tweet properties
const PropertyKeyTweetID string = "tweet_id"
const PropertyKeyUsername string = "username"
const PropertyKeyAuthorName string = "author_name"
const PropertyKeyRetweetCount string = "retweet_count"
const PropertyKeyAuthorID string = "author_id"
const PropertyKeyMediaURLs string = "media_urls"
const PropertyKeyMediaInfo string = "media_info"
const PropertyKeyExternalURLs string = "external_urls"
const PropertyKeyQuoteCount string = "quote_count"
const PropertyKeyConversationID string = "conversation_id"
const PropertyKeyLanguage string = "language"
const PropertyKeySource string = "source"
const PropertyKeyHasMedia string = "has_media"
const PropertyKeyMediaCount string = "media_count"
const PropertyKeyIsRetweet string = "is_retweet"
const PropertyKeyIsReply string = "is_reply"
const PropertyKeyIsQuote string = "is_quote"

// Email properties
const PropertyKeyEmailID string = "email_id"
const PropertyKeyThreadID string = "thread_id"
const PropertyKeyFrom string = "from"
const PropertyKeyTo string = "to"
const PropertyKeySubject string = "subject"
const PropertyKeyDate string = "date"
const PropertyKeyReceivedAt string = "received_at"
const PropertyKeyAttachments string = "attachments"
const PropertyKeyLabels string = "labels"

// Book properties
const PropertyKeyISBN string = "isbn"
const PropertyKeyPublisher string = "publisher"
const PropertyKeyPageCount string = "page_count"
const PropertyKeyPrice string = "price"
const PropertyKeyAuthorBio string = "author_bio"
const PropertyKeyTableOfContents string = "table_of_contents"
const PropertyKeyReviews string = "reviews"

// Person properties
const PropertyKeyFirstName string = "first_name"
const PropertyKeyLastName string = "last_name"
const PropertyKeyBirthday string = "birthday"
const PropertyKeyRelationType string = "relation_type"
const PropertyKeyAddress string = "address"
const PropertyKeyPhoneNumber string = "phone_number"

// Place properties
const PropertyKeyPlaceType string = "place_type"
const PropertyKeyCoordinates string = "coordinates"
const PropertyKeyMapURL string = "map_url"
const PropertyKeyVisitedDate string = "visited_date"
const PropertyKeyPlaceReviews string = "place_reviews"

// Media properties (for video, audio, image, file)
const PropertyKeySize string = "size"
const PropertyKeyTranscription string = "transcription"
const PropertyKeyPublicURL string = "public_url"
const PropertyKeyFilename string = "filename"
const PropertyKeyExtension string = "extension"

// propertyTypes maps property keys to their expected types
var propertyTypes = map[string]PropertyType{
	// Common properties
	PropertyKeyTitle:          TypeString,
	PropertyKeyText:           TypeString,
	PropertyKeyDescription:    TypeString,
	PropertyKeyTargetDateTime: TypeDateTime,
	PropertyKeyReminderOffset: TypeInt,
	PropertyKeyChecked:        TypeBool,
	PropertyKeyEnriched:       TypeBool,

	// URL/Link properties
	PropertyKeyURL:      TypeString,
	PropertyKeyImageURL: TypeString,

	// Movie/TV Series properties
	PropertyKeyIMDBID:           TypeString,
	PropertyKeyTMDBID:           TypeString,
	PropertyKeyReleaseYear:      TypeInt,
	PropertyKeyFirstAirYear:     TypeInt,
	PropertyKeyLastAirYear:      TypeInt,
	PropertyKeyRating:           TypeFloat,
	PropertyKeyRuntime:          TypeInt,
	PropertyKeyGenres:           TypeStringArray,
	PropertyKeyDirectors:        TypeStringArray,
	PropertyKeyCreators:         TypeStringArray,
	PropertyKeyCast:             TypeStringArray,
	PropertyKeyTagline:          TypeString,
	PropertyKeyBudget:           TypeInt,
	PropertyKeyRevenue:          TypeInt,
	PropertyKeyStatus:           TypeString,
	PropertyKeyInProduction:     TypeBool,
	PropertyKeyNumberOfSeasons:  TypeInt,
	PropertyKeyNumberOfEpisodes: TypeInt,
	PropertyKeyNetworks:         TypeStringArray,
	PropertyKeyType:             TypeString,

	// YouTube properties
	PropertyKeyVideoID:      TypeString,
	PropertyKeyChannelID:    TypeString,
	PropertyKeyChannelTitle: TypeString,
	PropertyKeyPublishedAt:  TypeDateTime,
	PropertyKeyViewCount:    TypeInt,
	PropertyKeyLikeCount:    TypeInt,
	PropertyKeyCommentCount: TypeInt,
	PropertyKeyDuration:     TypeString,
	PropertyKeyDefinition:   TypeString,
	PropertyKeyHasCaptions:  TypeBool,
	PropertyKeyTags:         TypeStringArray,

	// Tweet properties
	PropertyKeyTweetID:        TypeString,
	PropertyKeyUsername:       TypeString,
	PropertyKeyAuthorName:     TypeString,
	PropertyKeyRetweetCount:   TypeInt,
	PropertyKeyAuthorID:       TypeString,
	PropertyKeyMediaURLs:      TypeStringArray,
	PropertyKeyMediaInfo:      TypeStringArray,
	PropertyKeyExternalURLs:   TypeStringArray,
	PropertyKeyQuoteCount:     TypeInt,
	PropertyKeyConversationID: TypeString,
	PropertyKeyLanguage:       TypeString,
	PropertyKeySource:         TypeString,
	PropertyKeyHasMedia:       TypeBool,
	PropertyKeyMediaCount:     TypeInt,
	PropertyKeyIsRetweet:      TypeBool,
	PropertyKeyIsReply:        TypeBool,
	PropertyKeyIsQuote:        TypeBool,

	// Email properties
	PropertyKeyEmailID:     TypeString,
	PropertyKeyThreadID:    TypeString,
	PropertyKeyFrom:        TypeString,
	PropertyKeyTo:          TypeStringArray,
	PropertyKeySubject:     TypeString,
	PropertyKeyDate:        TypeDateTime,
	PropertyKeyReceivedAt:  TypeDateTime,
	PropertyKeyAttachments: TypeStringArray,
	PropertyKeyLabels:      TypeStringArray,

	// Book properties
	PropertyKeyISBN:            TypeString,
	PropertyKeyPublisher:       TypeString,
	PropertyKeyPageCount:       TypeInt,
	PropertyKeyPrice:           TypeFloat,
	PropertyKeyAuthorBio:       TypeString,
	PropertyKeyTableOfContents: TypeString,
	PropertyKeyReviews:         TypeStringArray,

	// Person properties
	PropertyKeyFirstName:    TypeString,
	PropertyKeyLastName:     TypeString,
	PropertyKeyBirthday:     TypeDateTime,
	PropertyKeyRelationType: TypeString,
	PropertyKeyAddress:      TypeString,
	PropertyKeyPhoneNumber:  TypeString,

	// Place properties
	PropertyKeyPlaceType:    TypeString,
	PropertyKeyCoordinates:  TypeFloatArray,
	PropertyKeyMapURL:       TypeString,
	PropertyKeyVisitedDate:  TypeDateTime,
	PropertyKeyPlaceReviews: TypeStringArray,

	// Media properties (for video, audio, image, file)
	PropertyKeySize:          TypeInt,
	PropertyKeyTranscription: TypeString,
	PropertyKeyPublicURL:     TypeString,
	PropertyKeyFilename:      TypeString,
	PropertyKeyExtension:     TypeString,
}

// getPropertyType returns the expected type for a property key
// If the key is not in the map, it returns TypeAny for backward compatibility
func getPropertyType(key string) PropertyType {
	if t, exists := propertyTypes[key]; exists {
		return t
	}
	return TypeAny
}

type Properties map[string][]interface{}

// IsEmpty checks if the properties map is empty or nil
func (p Properties) IsEmpty() bool {
	return p == nil || len(p) == 0
}

// MarshalJSON implements the json.Marshaler interface for Properties.
// It converts the Properties map to a valid JSON object.
func (p Properties) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string][]interface{}(p))
}

// UnmarshalJSON implements the json.Unmarshaler interface for Properties.
// It parses a JSON object into a Properties map.
func (p *Properties) UnmarshalJSON(data []byte) error {
	// Create a temporary map to unmarshal into
	var temp map[string][]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Initialize if nil
	if *p == nil {
		*p = Properties{}
	}

	// Copy values from temp map to Properties
	for k, v := range temp {
		(*p)[k] = v
	}

	return nil
}

// Has checks if a property exists and has at least one value
func (p Properties) Has(key string) bool {
	values, exists := p[key]
	return exists && len(values) > 0
}

// Get returns the first value for a given key and a boolean indicating whether the value was found
// This follows the idiomatic Go pattern of returning (value, ok)
func (p Properties) Get(key string) (interface{}, bool) {
	if values, exists := p[key]; exists && len(values) > 0 {
		return values[0], true
	}
	return nil, false
}

// GetString returns the first value as a string for a given key
func (p Properties) GetString(key string) (string, bool) {
	if val, ok := p.Get(key); ok {
		if str, ok := val.(string); ok {
			return str, true
		}
		// Try to convert to string
		return fmt.Sprintf("%v", val), true
	}
	return "", false
}

// GetInt returns the first value as an int for a given key
func (p Properties) GetInt(key string) (int, bool) {
	if val, ok := p.Get(key); ok {
		switch v := val.(type) {
		case int:
			return v, true
		case float64:
			return int(v), true
		case string:
			if i, err := strconv.Atoi(v); err == nil {
				return i, true
			}
		case time.Duration:
			return int(v.Seconds()), true
		}
	}
	return 0, false
}

// GetFloat returns the first value as a float64 for a given key
func (p Properties) GetFloat(key string) (float64, bool) {
	if val, ok := p.Get(key); ok {
		switch v := val.(type) {
		case float64:
			return v, true
		case int:
			return float64(v), true
		case string:
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				return f, true
			}
		}
	}
	return 0, false
}

// GetBool returns the first value as a bool for a given key
func (p Properties) GetBool(key string) (bool, bool) {
	if val, ok := p.Get(key); ok {
		switch v := val.(type) {
		case bool:
			return v, true
		case string:
			if b, err := strconv.ParseBool(v); err == nil {
				return b, true
			}
		case int:
			return v != 0, true
		case float64:
			return v != 0, true
		}
	}
	return false, false
}

// GetTime returns the first value as a time.Time for a given key
func (p Properties) GetTime(key string) (time.Time, bool) {
	if val, ok := p.Get(key); ok {
		switch v := val.(type) {
		case time.Time:
			return v, true
		case string:
			// Try to parse the time string - add more layouts as needed
			layouts := []string{
				time.RFC3339,
				"2006-01-02",
				"2006-01-02 15:04:05",
			}
			for _, layout := range layouts {
				if t, err := time.Parse(layout, v); err == nil {
					return t, true
				}
			}
		}
	}
	return time.Time{}, false
}

// GetStringArray returns all values as strings for a given key
func (p Properties) GetStringArray(key string) ([]string, bool) {
	if values, ok := p.GetArray(key); ok {
		result := make([]string, 0, len(values))
		for _, v := range values {
			if str, ok := v.(string); ok {
				result = append(result, str)
			} else {
				// Convert to string
				result = append(result, fmt.Sprintf("%v", v))
			}
		}
		return result, true
	}
	return nil, false
}

// GetArray returns all values for a given key and a boolean indicating whether the key exists
func (p Properties) GetArray(key string) ([]interface{}, bool) {
	if values, exists := p[key]; exists {
		return values, true
	}
	return nil, false
}

// convertValue converts a value to the expected type for a property key
func convertValue(key string, value interface{}) (interface{}, error) {
	propType := getPropertyType(key)

	switch propType {
	case TypeString:
		switch v := value.(type) {
		case string:
			return v, nil
		default:
			return fmt.Sprintf("%v", v), nil
		}

	case TypeInt:
		switch v := value.(type) {
		case int:
			return v, nil
		case float64:
			return int(v), nil
		case string:
			i, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("cannot convert %v to int: %w", v, err)
			}
			return i, nil
		case time.Duration:
			return int(v.Seconds()), nil
		default:
			return nil, fmt.Errorf("cannot convert %T to int", value)
		}

	case TypeFloat:
		switch v := value.(type) {
		case float64:
			return v, nil
		case int:
			return float64(v), nil
		case string:
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, fmt.Errorf("cannot convert %v to float: %w", v, err)
			}
			return f, nil
		default:
			return nil, fmt.Errorf("cannot convert %T to float", value)
		}

	case TypeBool:
		switch v := value.(type) {
		case bool:
			return v, nil
		case string:
			b, err := strconv.ParseBool(v)
			if err != nil {
				return nil, fmt.Errorf("cannot convert %v to bool: %w", v, err)
			}
			return b, nil
		case int:
			return v != 0, nil
		case float64:
			return v != 0, nil
		default:
			return nil, fmt.Errorf("cannot convert %T to bool", value)
		}

	case TypeDateTime:
		switch v := value.(type) {
		case time.Time:
			return v, nil
		case string:
			// Try to parse the time string - add more layouts as needed
			layouts := []string{
				time.RFC3339,
				"2006-01-02",
				"2006-01-02 15:04:05",
			}
			for _, layout := range layouts {
				if t, err := time.Parse(layout, v); err == nil {
					return t, nil
				}
			}
			return nil, fmt.Errorf("cannot parse %v as time", v)
		default:
			return nil, fmt.Errorf("cannot convert %T to time.Time", value)
		}

	case TypeStringArray:
		// If it's already an array, let it pass through
		// For single values, we'll handle them in the AppendValue/ReplaceValue methods
		return value, nil

	case TypeFloatArray:
		// If it's already an array, let it pass through
		// For single values, we'll handle them in the AppendValue/ReplaceValue methods
		return value, nil

	case TypeAny:
		// No conversion for backward compatibility
		return value, nil

	default:
		return value, nil
	}
}

// AppendValue safely appends a value to a property, creating the array if it doesn't exist yet
func (p Properties) AppendValue(key string, value interface{}) error {
	if p == nil {
		return fmt.Errorf("properties is nil")
	}

	// Convert value to the expected type
	convertedValue, err := convertValue(key, value)
	if err != nil {
		return fmt.Errorf("failed to append value: %w", err)
	}

	if _, exists := p[key]; !exists {
		p[key] = []interface{}{}
	}

	p[key] = append(p[key], convertedValue)
	return nil
}

// ReplaceValue replaces value of the key
func (p Properties) ReplaceValue(key string, value interface{}) error {
	if p == nil {
		return fmt.Errorf("properties is nil")
	}

	// Convert value to the expected type
	convertedValue, err := convertValue(key, value)
	if err != nil {
		return fmt.Errorf("failed to replace value: %w", err)
	}

	// Reset and append
	p[key] = []interface{}{}
	p[key] = append(p[key], convertedValue)
	return nil
}

// Delete removes a property with the given key and returns a boolean
// indicating whether the property existed
func (p Properties) Delete(key string) bool {
	if p == nil {
		return false
	}

	_, exists := p[key]
	if exists {
		delete(p, key)
	}
	return exists
}
