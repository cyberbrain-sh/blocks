package blocks

import (
	"fmt"
)

func AddBulletListItemProperties(b *Block, title *string) error {
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

func RenderBulletListItemProperties(b Block) string {
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

func GetBulletListItemProperties() []string {
	return []string{
		PropertyKeyTitle,
	}
}
