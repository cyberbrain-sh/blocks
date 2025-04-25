package blocks

import (
	"fmt"
)

func AddHeaderProperties(b *Block, title *string) error {
	if b == nil {
		return fmt.Errorf("cannot add header properties because given b is nil")
	}

	if title != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTitle, *title); err != nil {
			return fmt.Errorf("failed to set title property: %w", err)
		}
	}

	return nil
}

func RenderHeaderProperties(b Block) string {
	text, ok := b.Properties.Get(PropertyKeyTitle)
	if !ok {
		text = ""
	}

	textString, ok := text.(string)
	if !ok {
		textString = ""
	}

	// Add # based on the header size in type
	var prefix string
	switch b.Type {
	case TypeHeader1:
		prefix = "# "
	case TypeHeader2:
		prefix = "## "
	case TypeHeader3:
		prefix = "### "
	case TypeHeader4:
		prefix = "#### "
	case TypeHeader5:
		prefix = "##### "
	case TypeHeader6:
		prefix = "###### "
	default:
		// For non-header types, don't add any prefix
		prefix = ""
	}

	return prefix + textString
}

func GetHeaderProperties() []string {
	return []string{
		PropertyKeyTitle,
	}
}
