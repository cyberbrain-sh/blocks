package pkg

import (
	"fmt"
)

func AddNumberedListItemProperties(b *Block, title *string) error {
	if b == nil {
		return fmt.Errorf("cannot add bullet list item properties because given b is nil")
	}

	if title != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTitle, *title); err != nil {
			return fmt.Errorf("failed to set title property: %w", err)
		}
	}

	return nil
}

func RenderNumberedListItemProperties(b Block) string {
	text, ok := b.Properties.Get(PropertyKeyTitle)
	if !ok {
		text = ""
	}

	textString, ok := text.(string)
	if !ok {
		textString = ""
	}

	return textString
}

func GetNumberedListItemProperties() []string {
	return []string{
		PropertyKeyTitle,
	}
}
