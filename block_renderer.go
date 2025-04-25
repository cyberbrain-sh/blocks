package blocks

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RenderContent(ctx context.Context, b Block, lookupBlocks map[uuid.UUID]Block) (string, error) {
	// Always render the current block's properties
	ownContent := RenderProperties(ctx, b)

	// If content is empty, return just the properties
	if len(b.Content) == 0 {
		return ownContent, nil
	}

	// Build the rendered content by processing each child block
	var renderedContent []string

	// Add own content if not empty
	if ownContent != "" {
		renderedContent = append(renderedContent, ownContent)
	}

	for _, blockID := range b.Content {
		// Find the child block in the lookup blocks using direct map access
		childBlock, exists := lookupBlocks[blockID]
		if !exists {
			// Skip if child block not found
			continue
		}

		// Recursively render the child block
		childContent, err := RenderContent(ctx, childBlock, lookupBlocks)
		if err != nil {
			return "", err
		}

		// Add the rendered content if not empty
		if childContent != "" {
			renderedContent = append(renderedContent, childContent)
		}
	}

	// Join all rendered content with newlines
	return strings.Join(renderedContent, "\n"), nil
}

type RenderingForJSONStructure struct {
	AnnotationID string                      `json:"annotation_id"`
	ID           string                      `json:"id"`
	Properties   json.RawMessage             `json:"properties"`
	ContentType  string                      `json:"content_type"`
	ChildBlocks  []RenderingForJSONStructure `json:"child_blocks"`
	Keywords     []string                    `json:"keywords"`
	CreatedAt    string                      `json:"created_at"`
	UpdatedAt    string                      `json:"updated_at"`
	Score        float64                     `json:"-"` // Avoid this to reduce confusion in reranker
}

func RenderAsJSON(ctx context.Context, b Block, lookupBlocks map[uuid.UUID]Block) RenderingForJSONStructure {

	// Recursive function to convert a block to RenderingForJSONStructure
	var buildRenderingStructure func(block Block) RenderingForJSONStructure
	buildRenderingStructure = func(block Block) RenderingForJSONStructure {
		// Create the base structure for this block
		propJson, _ := json.Marshal(block.Properties)

		result := RenderingForJSONStructure{
			AnnotationID: block.AnnotationID(),
			ID:           block.ID.String(),
			Properties:   propJson,
			ContentType:  block.Type.ContentType(),
			Keywords:     block.Keywords(),
			CreatedAt:    block.CreatedAt.Format(time.RFC3339),
			UpdatedAt:    block.UpdatedAt.Format(time.RFC3339),
		}

		// Process child blocks
		for _, childID := range block.Content {
			// Find the child block in the lookup using direct map access
			childBlock, exists := lookupBlocks[childID]
			if !exists {
				// Skip if child block not found
				continue
			}

			// Recursively process the child block
			childStructure := buildRenderingStructure(childBlock)
			result.ChildBlocks = append(result.ChildBlocks, childStructure)
		}

		return result
	}

	// Build the complete structure starting from the root block
	rootStructure := buildRenderingStructure(b)

	return rootStructure
}
