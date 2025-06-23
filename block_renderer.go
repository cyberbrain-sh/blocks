package blocks

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RenderContent(ctx context.Context, b Block, lookupBlocks map[uuid.UUID]Block) (string, error) {
	// Create a new context with visited blocks map if it doesn't exist
	return renderContentWithCycleDetection(ctx, b, lookupBlocks, make(map[uuid.UUID]bool))
}

// renderContentWithCycleDetection is an internal function that tracks visited blocks to prevent infinite recursion
func renderContentWithCycleDetection(ctx context.Context, b Block, lookupBlocks map[uuid.UUID]Block, visitedBlocks map[uuid.UUID]bool) (string, error) {
	// Check if we've already processed this block to prevent cycles
	if visitedBlocks[b.ID] {
		// We detected a cycle, return empty string to break the recursion
		return "", nil
	}

	// Mark this block as visited
	visitedBlocks[b.ID] = true

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

		// Recursively render the child block with cycle detection
		childContent, err := renderContentWithCycleDetection(ctx, childBlock, lookupBlocks, visitedBlocks)
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
	AnnotationID string                       `json:"annotation_id"`
	ID           string                       `json:"id"`
	Properties   json.RawMessage              `json:"properties"`
	ContentType  string                       `json:"content_type"`
	ChildBlocks  []RenderingForJSONStructure  `json:"child_blocks"`
	Keywords     []string                     `json:"keywords"`
	CreatedAt    string                       `json:"created_at"`
	UpdatedAt    string                       `json:"updated_at"`
	Space        *RenderingSpaceJSONStructure `json:"space,omitempty"`
	Score        float64                      `json:"-"` // Avoid this to reduce confusion in reranker
}

type RenderingSpaceJSONStructure struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func RenderAsJSON(ctx context.Context, b Block, lookupBlocks map[uuid.UUID]Block) RenderingForJSONStructure {
	// Apply the same cycle detection to RenderAsJSON
	return renderAsJSONWithCycleDetection(b, lookupBlocks, make(map[uuid.UUID]bool))
}

// renderAsJSONWithCycleDetection is an internal function that prevents infinite recursion in RenderAsJSON
func renderAsJSONWithCycleDetection(block Block, lookupBlocks map[uuid.UUID]Block, visitedBlocks map[uuid.UUID]bool) RenderingForJSONStructure {
	// Create the base structure for this block
	propJson, _ := json.Marshal(block.Properties)

	result := RenderingForJSONStructure{
		AnnotationID: block.AnnotationID(),
		ID:           block.ID.String(),
		Properties:   propJson,
		ContentType:  block.Type.ContentType().String(),
		Keywords:     block.Keywords(),
		CreatedAt:    block.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    block.UpdatedAt.Format(time.RFC3339),
	}

	// Process child blocks
	for _, childID := range block.Content {
		// Skip if we've already processed this block (cycle detection)
		if visitedBlocks[childID] {
			continue
		}

		// Mark this block as visited
		visitedBlocks[childID] = true

		// Find the child block in the lookup using direct map access
		childBlock, exists := lookupBlocks[childID]
		if !exists {
			// Skip if child block not found
			continue
		}

		// Recursively process the child block with cycle detection
		childStructure := renderAsJSONWithCycleDetection(childBlock, lookupBlocks, visitedBlocks)
		result.ChildBlocks = append(result.ChildBlocks, childStructure)
	}

	return result
}
