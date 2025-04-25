package blocks

import (
	"fmt"
)

func AddLinkPropertiesFromPage(block *Block, page PageData) error {
	if block == nil {
		return fmt.Errorf("cannot add link properties because given block is nil")
	}

	// Extract basic page information
	title := page.GetTitle()
	description := page.GetDescription()
	imageURL := page.GetImage()
	originalURL := page.GetURL()

	// Call AddLinkProperties with the extracted values
	return AddLinkProperties(block, &originalURL, &title, &description, &imageURL, true)
}

func AddLinkProperties(b *Block, url, title, description, urlImage *string, enriched bool) error {
	if b == nil {
		return fmt.Errorf("cannot add link properties because given b is nil")
	}

	if url != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyURL, *url); err != nil {
			return fmt.Errorf("failed to set URL property: %w", err)
		}
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

	if urlImage != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyImageURL, *urlImage); err != nil {
			return fmt.Errorf("failed to set image URL property: %w", err)
		}
	}

	if err := b.Properties.ReplaceValue(PropertyKeyEnriched, enriched); err != nil {
		return fmt.Errorf("failed to set enriched property: %w", err)
	}

	return nil
}

func RenderLinkProperties(b Block) string {
	// Get URL
	urlValue, hasURL := b.Properties.Get(PropertyKeyURL)
	url, urlOk := urlValue.(string)

	// Get title
	titleValue, hasTitle := b.Properties.Get(PropertyKeyTitle)
	title, titleOk := titleValue.(string)

	// Get description
	descValue, hasDesc := b.Properties.Get(PropertyKeyDescription)
	description, descOk := descValue.(string)

	// Format as markdown link if we have both URL and title
	if hasURL && urlOk && hasTitle && titleOk && url != "" && title != "" {
		if hasDesc && descOk && description != "" {
			// Include description if available
			return fmt.Sprintf("[%s](%s) - %s", title, url, description)
		}
		// Just URL and title
		return fmt.Sprintf("[%s](%s)", title, url)
	} else if hasURL && urlOk && url != "" {
		// Just URL if no title
		return url
	} else if hasTitle && titleOk && title != "" {
		// Just title if no URL
		return title
	}

	return ""
}

func GetLinkProperties() []string {
	return []string{
		PropertyKeyTitle,
		PropertyKeyURL,
		PropertyKeyDescription,
		PropertyKeyImageURL,
		PropertyKeyEnriched,
	}
}
