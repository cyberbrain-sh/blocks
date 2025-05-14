package blocks

import (
	"fmt"
)

// AddLineProperties is an empty function since line blocks don't have any properties
func AddLineProperties(b *Block) error {
	if b == nil {
		return fmt.Errorf("cannot add line properties because given block is nil")
	}

	// Line blocks have no properties to set

	return nil
}

// RenderLineProperties renders a line block as a markdown horizontal rule
func RenderLineProperties(b Block) string {
	// Always renders as a horizontal rule in markdown
	return "---"
}

// GetLineProperties returns an empty list since line blocks have no properties
func GetLineProperties() []string {
	return []string{}
}
