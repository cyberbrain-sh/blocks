package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRenderContent(t *testing.T) {
	// Setup common test context
	ctx := context.Background()

	t.Run("empty content should render properties", func(t *testing.T) {
		// Create a block with empty content but with properties
		block := Block{
			ID:         uuid.New(),
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Test Title"}},
			Content:    []uuid.UUID{}, // Empty content
		}

		// Call RenderContent
		result, err := RenderContent(ctx, block, nil)

		// Assert no error
		assert.NoError(t, err)
		// Assert result matches the title property rendered by RenderParagraphProperties
		assert.Equal(t, "Test Title", result)
	})

	t.Run("single child block", func(t *testing.T) {
		// Create parent and child blocks
		childID := uuid.New()
		childBlock := Block{
			ID:         childID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Child Block Content"}},
			Content:    []uuid.UUID{}, // Child has no further content
		}

		parentBlock := Block{
			ID:         uuid.New(),
			Type:       TypePage,
			Properties: Properties{PropertyKeyTitle: {"Parent Block Content"}},
			Content:    []uuid.UUID{childID}, // Parent references child
		}

		// Create a map of blocks for lookupBlocks
		blocksMap := map[uuid.UUID]Block{
			childID: childBlock,
		}

		// Call RenderContent with blocks in lookupBlocks map
		result, err := RenderContent(ctx, parentBlock, blocksMap)

		// Assert no error
		assert.NoError(t, err)
		// Now parent block's properties should be rendered along with the child's
		assert.Equal(t, "Parent Block Content\nChild Block Content", result)
	})

	t.Run("missing child block", func(t *testing.T) {
		// Create parent block referencing a non-existent child
		nonExistentID := uuid.New()
		parentBlock := Block{
			ID:         uuid.New(),
			Type:       TypePage,
			Properties: Properties{PropertyKeyTitle: {"Parent Block Content"}},
			Content:    []uuid.UUID{nonExistentID}, // References a non-existent block
		}

		// Call RenderContent with empty lookupBlocks map
		result, err := RenderContent(ctx, parentBlock, map[uuid.UUID]Block{})

		// Assert no error
		assert.NoError(t, err)
		// Now parent block's properties should be rendered even if child is missing
		assert.Equal(t, "Parent Block Content", result)
	})

	t.Run("multiple child blocks", func(t *testing.T) {
		// Create multiple child blocks
		child1ID := uuid.New()
		child1Block := Block{
			ID:         child1ID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Child 1 Content"}},
			Content:    []uuid.UUID{},
		}

		child2ID := uuid.New()
		child2Block := Block{
			ID:         child2ID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Child 2 Content"}},
			Content:    []uuid.UUID{},
		}

		parentBlock := Block{
			ID:         uuid.New(),
			Type:       TypePage,
			Properties: Properties{PropertyKeyTitle: {"Parent Block Content"}},
			Content:    []uuid.UUID{child1ID, child2ID}, // References both children
		}

		// Create a map of blocks for lookupBlocks
		blocksMap := map[uuid.UUID]Block{
			child1ID: child1Block,
			child2ID: child2Block,
		}

		// Call RenderContent with all blocks in lookupBlocks map
		result, err := RenderContent(ctx, parentBlock, blocksMap)

		// Assert no error
		assert.NoError(t, err)
		// Parent and both children's properties should be rendered
		assert.Equal(t, "Parent Block Content\nChild 1 Content\nChild 2 Content", result)
	})

	t.Run("recursive rendering", func(t *testing.T) {
		// Create a nested structure of blocks
		grandchildID := uuid.New()
		grandchildBlock := Block{
			ID:         grandchildID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Grandchild Content"}},
			Content:    []uuid.UUID{},
		}

		childID := uuid.New()
		childBlock := Block{
			ID:         childID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Child Content"}}, // Now child has title
			Content:    []uuid.UUID{grandchildID},                       // Child references grandchild
		}

		parentBlock := Block{
			ID:         uuid.New(),
			Type:       TypePage,
			Properties: Properties{PropertyKeyTitle: {"Parent Content"}},
			Content:    []uuid.UUID{childID}, // Parent references child
		}

		// Create a map of blocks for lookupBlocks
		blocksMap := map[uuid.UUID]Block{
			childID:      childBlock,
			grandchildID: grandchildBlock,
		}

		// Call RenderContent with all blocks in lookupBlocks map
		result, err := RenderContent(ctx, parentBlock, blocksMap)

		// Assert no error
		assert.NoError(t, err)
		// Parent, child, and grandchild properties should all be rendered
		assert.Equal(t, "Parent Content\nChild Content\nGrandchild Content", result)
	})

	t.Run("child with both properties and content", func(t *testing.T) {
		// Create a situation where a child block has both properties and content
		grandchildID := uuid.New()
		grandchildBlock := Block{
			ID:         grandchildID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Grandchild Content"}},
			Content:    []uuid.UUID{},
		}

		childID := uuid.New()
		childBlock := Block{
			ID:         childID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Child Content"}}, // Non-empty title
			Content:    []uuid.UUID{grandchildID},                       // Also has content
		}

		parentBlock := Block{
			ID:         uuid.New(),
			Type:       TypePage,
			Properties: Properties{PropertyKeyTitle: {"Parent Content"}},
			Content:    []uuid.UUID{childID},
		}

		// Create a map of blocks for lookupBlocks
		blocksMap := map[uuid.UUID]Block{
			childID:      childBlock,
			grandchildID: grandchildBlock,
		}

		// Call RenderContent with all blocks in lookupBlocks map
		result, err := RenderContent(ctx, parentBlock, blocksMap)

		// Assert no error
		assert.NoError(t, err)
		// Now all blocks' properties should be rendered
		assert.Equal(t, "Parent Content\nChild Content\nGrandchild Content", result)
	})

	t.Run("multiple levels with multiple children", func(t *testing.T) {
		// Create multiple levels with multiple children to test newline rendering
		grandchild1ID := uuid.New()
		grandchild1Block := Block{
			ID:         grandchild1ID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Grandchild 1 Content"}},
			Content:    []uuid.UUID{},
		}

		grandchild2ID := uuid.New()
		grandchild2Block := Block{
			ID:         grandchild2ID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Grandchild 2 Content"}},
			Content:    []uuid.UUID{},
		}

		child1ID := uuid.New()
		child1Block := Block{
			ID:         child1ID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Child 1 Content"}},
			Content:    []uuid.UUID{grandchild1ID, grandchild2ID}, // Child 1 has two grandchildren
		}

		child2ID := uuid.New()
		child2Block := Block{
			ID:         child2ID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Child 2 Content"}},
			Content:    []uuid.UUID{}, // Child 2 has no children
		}

		parentBlock := Block{
			ID:         uuid.New(),
			Type:       TypePage,
			Properties: Properties{PropertyKeyTitle: {"Parent Content"}},
			Content:    []uuid.UUID{child1ID, child2ID}, // Parent has two children
		}

		// Create a map of blocks for lookupBlocks
		blocksMap := map[uuid.UUID]Block{
			child1ID:      child1Block,
			child2ID:      child2Block,
			grandchild1ID: grandchild1Block,
			grandchild2ID: grandchild2Block,
		}

		// Call RenderContent with all blocks in lookupBlocks map
		result, err := RenderContent(ctx, parentBlock, blocksMap)

		// Assert no error
		assert.NoError(t, err)
		// All blocks' properties should be rendered
		expectedResult := "Parent Content\nChild 1 Content\nGrandchild 1 Content\nGrandchild 2 Content\nChild 2 Content"
		assert.Equal(t, expectedResult, result)
	})

	t.Run("empty properties should not add empty lines", func(t *testing.T) {
		// Create a block with empty properties
		childID := uuid.New()
		childBlock := Block{
			ID:         childID,
			Type:       TypeParagraph,
			Properties: Properties{PropertyKeyTitle: {"Child Content"}},
			Content:    []uuid.UUID{},
		}

		parentBlock := Block{
			ID:         uuid.New(),
			Type:       TypePage,
			Properties: Properties{PropertyKeyTitle: {""}}, // Empty title
			Content:    []uuid.UUID{childID},
		}

		// Create a map of blocks for lookupBlocks
		blocksMap := map[uuid.UUID]Block{
			childID: childBlock,
		}

		// Call RenderContent with blocks in lookupBlocks map
		result, err := RenderContent(ctx, parentBlock, blocksMap)

		// Assert no error
		assert.NoError(t, err)
		// Only non-empty properties should be rendered
		assert.Equal(t, "Child Content", result)
	})
}

func TestRenderAsJSON(t *testing.T) {
	ctx := context.Background()

	// Helper function to parse and compare JSON
	compareJSON := func(t *testing.T, expected string, actual RenderingForJSONStructure) {
		t.Helper()
		var expectedObj map[string]interface{}
		err := json.Unmarshal([]byte(expected), &expectedObj)
		assert.NoError(t, err, "Failed to parse expected JSON")

		// Convert actual to JSON string first
		actualBytes, err := json.Marshal(actual)
		assert.NoError(t, err, "Failed to marshal actual to JSON")

		var actualObj map[string]interface{}
		err = json.Unmarshal(actualBytes, &actualObj)
		assert.NoError(t, err, "Failed to parse actual JSON")

		// Improved function to normalize nil values to empty arrays
		var normalizeNilToEmptyArray func(interface{}) interface{}
		normalizeNilToEmptyArray = func(value interface{}) interface{} {
			switch v := value.(type) {
			case map[string]interface{}:
				result := make(map[string]interface{})
				for key, val := range v {
					// For child_blocks and keywords fields, always ensure they are arrays
					if key == "child_blocks" || key == "keywords" {
						if val == nil {
							result[key] = []interface{}{}
						} else if array, ok := val.([]interface{}); ok {
							// Recursively normalize each element in the array
							normalized := make([]interface{}, len(array))
							for i, elem := range array {
								normalized[i] = normalizeNilToEmptyArray(elem)
							}
							result[key] = normalized
						} else {
							// If it's not nil and not an array, keep it as is (shouldn't happen)
							result[key] = val
						}
					} else {
						// For other fields, recursively normalize
						result[key] = normalizeNilToEmptyArray(val)
					}
				}
				return result
			case []interface{}:
				// Normalize each element in the array
				result := make([]interface{}, len(v))
				for i, elem := range v {
					result[i] = normalizeNilToEmptyArray(elem)
				}
				return result
			default:
				// For all other types, return as is
				return v
			}
		}

		// Apply normalization to both objects
		expectedObj = normalizeNilToEmptyArray(expectedObj).(map[string]interface{})
		actualObj = normalizeNilToEmptyArray(actualObj).(map[string]interface{})

		assert.Equal(t, expectedObj, actualObj, "JSON structures should match")
	}

	t.Run("SingleBlockNoChildren", func(t *testing.T) {
		// Setup a single block with no children
		now := time.Now()
		blockID := uuid.New()

		props := Properties{}
		props.AppendValue(PropertyKeyTitle, "Test Block")

		block := Block{
			ID:         blockID,
			Type:       TypeParagraph,
			Properties: props,
			Content:    []uuid.UUID{},
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		// Call RenderAsJSON with empty map
		result := RenderAsJSON(ctx, block, map[uuid.UUID]Block{})

		// Expected JSON structure
		expected := fmt.Sprintf(`{
			"annotation_id": "[[block:%s]]",
			"id": "%s",
			"properties": {"title":["Test Block"]},
			"content_type": "textual",
			"child_blocks": [],
			"keywords": [],
			"created_at": "%s",
			"updated_at": "%s"
		}`, blockID, blockID, now.Format(time.RFC3339), now.Format(time.RFC3339))

		compareJSON(t, expected, result)
	})

	t.Run("BlockWithChildren", func(t *testing.T) {
		// Setup a block with children
		now := time.Now()
		parentID := uuid.New()
		child1ID := uuid.New()
		child2ID := uuid.New()

		parentProps := Properties{}
		parentProps.AppendValue(PropertyKeyTitle, "Parent Block")

		child1Props := Properties{}
		child1Props.AppendValue(PropertyKeyText, "Child 1 Text")

		child2Props := Properties{}
		child2Props.AppendValue(PropertyKeyText, "Child 2 Text")

		parent := Block{
			ID:         parentID,
			Type:       TypePage,
			Properties: parentProps,
			Content:    []uuid.UUID{child1ID, child2ID},
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		child1 := Block{
			ID:         child1ID,
			Type:       TypeParagraph,
			Properties: child1Props,
			Content:    []uuid.UUID{},
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		child2 := Block{
			ID:         child2ID,
			Type:       TypeParagraph,
			Properties: child2Props,
			Content:    []uuid.UUID{},
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		// Create a map of blocks instead of a slice
		blocksMap := map[uuid.UUID]Block{
			parentID: parent,
			child1ID: child1,
			child2ID: child2,
		}

		// Call RenderAsJSON with the map
		result := RenderAsJSON(ctx, parent, blocksMap)

		// Expected JSON structure
		expected := fmt.Sprintf(`{
			"annotation_id": "[[block:%s]]",
			"id": "%s",
			"properties": {"title":["Parent Block"]},
			"content_type": "textual",
			"child_blocks": [
				{
					"annotation_id": "[[block:%s]]",
					"id": "%s",
					"properties": {"text":["Child 1 Text"]},
					"content_type": "textual",
					"child_blocks": [],
					"keywords": [],
					"created_at": "%s",
					"updated_at": "%s"
				},
				{
					"annotation_id": "[[block:%s]]",
					"id": "%s",
					"properties": {"text":["Child 2 Text"]},
					"content_type": "textual",
					"child_blocks": [],
					"keywords": [],
					"created_at": "%s",
					"updated_at": "%s"
				}
			],
			"keywords": [],
			"created_at": "%s",
			"updated_at": "%s"
		}`,
			parentID, parentID,
			child1ID, child1ID, now.Format(time.RFC3339), now.Format(time.RFC3339),
			child2ID, child2ID, now.Format(time.RFC3339), now.Format(time.RFC3339),
			now.Format(time.RFC3339), now.Format(time.RFC3339))

		compareJSON(t, expected, result)
	})

	t.Run("NestedBlocks", func(t *testing.T) {
		// Setup a block with nested children (grandchildren)
		now := time.Now()
		parentID := uuid.New()
		childID := uuid.New()
		grandchildID := uuid.New()

		parentProps := Properties{}
		parentProps.AppendValue(PropertyKeyTitle, "Parent Block")

		childProps := Properties{}
		childProps.AppendValue(PropertyKeyText, "Child Text")

		grandchildProps := Properties{}
		grandchildProps.AppendValue(PropertyKeyText, "Grandchild Text")

		parent := Block{
			ID:         parentID,
			Type:       TypePage,
			Properties: parentProps,
			Content:    []uuid.UUID{childID},
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		child := Block{
			ID:         childID,
			Type:       TypeParagraph,
			Properties: childProps,
			Content:    []uuid.UUID{grandchildID},
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		grandchild := Block{
			ID:         grandchildID,
			Type:       TypeParagraph,
			Properties: grandchildProps,
			Content:    []uuid.UUID{},
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		// Create a map of blocks instead of a slice
		blocksMap := map[uuid.UUID]Block{
			parentID:     parent,
			childID:      child,
			grandchildID: grandchild,
		}

		// Call RenderAsJSON with the map
		result := RenderAsJSON(ctx, parent, blocksMap)

		// Expected JSON structure
		expected := fmt.Sprintf(`{
			"annotation_id": "[[block:%s]]",
			"id": "%s",
			"properties": {"title":["Parent Block"]},
			"content_type": "textual",
			"child_blocks": [
				{
					"annotation_id": "[[block:%s]]",
					"id": "%s",
					"properties": {"text":["Child Text"]},
					"content_type": "textual",
					"child_blocks": [
						{
							"annotation_id": "[[block:%s]]",
							"id": "%s",
							"properties": {"text":["Grandchild Text"]},
							"content_type": "textual",
							"child_blocks": [],
							"keywords": [],
							"created_at": "%s",
							"updated_at": "%s"
						}
					],
					"keywords": [],
					"created_at": "%s",
					"updated_at": "%s"
				}
			],
			"keywords": [],
			"created_at": "%s",
			"updated_at": "%s"
		}`,
			parentID, parentID,
			childID, childID,
			grandchildID, grandchildID, now.Format(time.RFC3339), now.Format(time.RFC3339),
			now.Format(time.RFC3339), now.Format(time.RFC3339),
			now.Format(time.RFC3339), now.Format(time.RFC3339))

		compareJSON(t, expected, result)
	})

	t.Run("MissingChildren", func(t *testing.T) {
		// Setup a block with references to children that don't exist in the lookup
		now := time.Now()
		parentID := uuid.New()
		missingChildID := uuid.New() // This child will be missing
		existingChildID := uuid.New()

		parentProps := Properties{}
		parentProps.AppendValue(PropertyKeyTitle, "Parent Block")

		childProps := Properties{}
		childProps.AppendValue(PropertyKeyText, "Existing Child")

		parent := Block{
			ID:         parentID,
			Type:       TypePage,
			Properties: parentProps,
			Content:    []uuid.UUID{missingChildID, existingChildID},
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		existingChild := Block{
			ID:         existingChildID,
			Type:       TypeParagraph,
			Properties: childProps,
			Content:    []uuid.UUID{},
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		// Create a map of blocks instead of a slice (note: missingChild is not included)
		blocksMap := map[uuid.UUID]Block{
			parentID:        parent,
			existingChildID: existingChild,
		}

		// Call RenderAsJSON with the map
		result := RenderAsJSON(ctx, parent, blocksMap)

		// Expected JSON structure (should not include the missing child)
		expected := fmt.Sprintf(`{
			"annotation_id": "[[block:%s]]",
			"id": "%s",
			"properties": {"title":["Parent Block"]},
			"content_type": "textual",
			"child_blocks": [
				{
					"annotation_id": "[[block:%s]]",
					"id": "%s",
					"properties": {"text":["Existing Child"]},
					"content_type": "textual",
					"child_blocks": [],
					"keywords": [],
					"created_at": "%s",
					"updated_at": "%s"
				}
			],
			"keywords": [],
			"created_at": "%s",
			"updated_at": "%s"
		}`,
			parentID, parentID,
			existingChildID, existingChildID, now.Format(time.RFC3339), now.Format(time.RFC3339),
			now.Format(time.RFC3339), now.Format(time.RFC3339))

		compareJSON(t, expected, result)
	})

	t.Run("EmptyProperties", func(t *testing.T) {
		// Setup a block with empty properties
		now := time.Now()
		blockID := uuid.New()

		block := Block{
			ID:         blockID,
			Type:       TypeParagraph,
			Properties: Properties{},
			Content:    []uuid.UUID{},
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		// Call RenderAsJSON with empty map
		result := RenderAsJSON(ctx, block, map[uuid.UUID]Block{})

		// Expected JSON structure
		expected := fmt.Sprintf(`{
			"annotation_id": "[[block:%s]]",
			"id": "%s",
			"properties": {},
			"content_type": "textual",
			"child_blocks": [],
			"keywords": [],
			"created_at": "%s",
			"updated_at": "%s"
		}`, blockID, blockID, now.Format(time.RFC3339), now.Format(time.RFC3339))

		compareJSON(t, expected, result)
	})
}
