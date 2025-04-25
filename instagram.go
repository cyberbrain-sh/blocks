package blocks

import (
	"fmt"
	"strings"
)

func AddInstagramPropertiesFromPage(block *Block, page PageData) error {
	if block == nil {
		return fmt.Errorf("cannot add instagram properties because given block is nil")
	}

	// Extract basic page information
	title := page.GetTitle()
	description := page.GetDescription()
	imageURL := page.GetImage()
	originalURL := page.GetURL()

	// Call AddInstagramProperties with the extracted values
	return AddInstagramProperties(block, &title, &description, &imageURL, &originalURL, true)
}

func AddInstagramProperties(b *Block, title, description, image, originalURL *string, enriched bool) error {
	if b == nil {
		return fmt.Errorf("cannot add instagram properties because given b is nil")
	}

	if title != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTitle, *title); err != nil {
			return fmt.Errorf("failed to set title property: %w", err)
		}
	}

	if description != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyDescription, *description); err != nil {
			return fmt.Errorf("failed to set description property: %w", err)
		}
	}

	if image != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyImageURL, *image); err != nil {
			return fmt.Errorf("failed to set image URL property: %w", err)
		}
	}

	if originalURL != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyURL, *originalURL); err != nil {
			return fmt.Errorf("failed to set URL property: %w", err)
		}
	}

	if err := b.Properties.ReplaceValue(PropertyKeyEnriched, enriched); err != nil {
		return fmt.Errorf("failed to set enriched property: %w", err)
	}

	return nil
}

func RenderInstagramProperties(b Block) string {
	// Get properties
	titleVal, hasTitle := b.Properties.Get(PropertyKeyTitle)
	title, titleOk := titleVal.(string)

	urlVal, hasURL := b.Properties.Get(PropertyKeyURL)
	url, urlOk := urlVal.(string)

	imageVal, hasImage := b.Properties.Get(PropertyKeyImageURL)
	image, imageOk := imageVal.(string)

	descVal, hasDesc := b.Properties.Get(PropertyKeyDescription)
	desc, descOk := descVal.(string)

	// Build the markdown representation
	var parts []string

	// Title as header, possibly with link
	if hasTitle && titleOk && title != "" {
		if hasURL && urlOk && url != "" {
			parts = append(parts, fmt.Sprintf("## [%s](%s)", title, url))
		} else {
			parts = append(parts, fmt.Sprintf("## %s", title))
		}
	} else if hasURL && urlOk && url != "" {
		parts = append(parts, fmt.Sprintf("## [Instagram Post](%s)", url))
	}

	// Image if available
	if hasImage && imageOk && image != "" {
		if hasURL && urlOk && url != "" {
			// Link the image to the original post
			parts = append(parts, fmt.Sprintf("[![Instagram Image](%s)](%s)", image, url))
		} else {
			parts = append(parts, fmt.Sprintf("![Instagram Image](%s)", image))
		}
	}

	// Description
	if hasDesc && descOk && desc != "" {
		parts = append(parts, desc)
	}

	return strings.Join(parts, "\n\n")
}

func GetInstagramProperties() []string {
	return []string{
		PropertyKeyTitle,
		PropertyKeyDescription,
		PropertyKeyImageURL,
		PropertyKeyURL,
		PropertyKeyEnriched,
	}
}
