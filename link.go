package blocks

import (
	"encoding/json"
	"fmt"
)

// LinkData represents the structure of link data in JSON format
type LinkData struct {
	URL         string `json:"url"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
}

func AddLinkPropertiesFromPage(block *Block, page PageData) error {
	if block == nil {
		return fmt.Errorf("cannot add link properties because given block is nil")
	}

	// Extract basic page information
	title := page.GetTitle()
	description := page.GetDescription()
	imageURL := page.GetImage()
	originalURL := page.GetURL()

	// Validate required URL field
	if originalURL == "" {
		return fmt.Errorf("page must include a URL")
	}

	// Call AddLinkProperties with the extracted values
	// AddLinkProperties will handle empty string checks for other fields
	return AddLinkProperties(block, &originalURL, &title, &description, &imageURL, true)
}

// AddLinkPropertiesFromJSON parses a json.RawMessage into a LinkData struct
// and adds the link properties to the given block
func AddLinkPropertiesFromJSON(block *Block, rawJSON json.RawMessage) error {
	if block == nil {
		return fmt.Errorf("cannot add link properties because given block is nil")
	}

	// Parse the JSON data into a LinkData struct
	var linkData LinkData
	if err := json.Unmarshal(rawJSON, &linkData); err != nil {
		return fmt.Errorf("failed to unmarshal link data: %w", err)
	}

	// Validate required fields
	if linkData.URL == "" {
		return fmt.Errorf("link data must include a URL")
	}

	// Convert struct fields to pointers for AddLinkProperties
	url := linkData.URL

	var title, description, imageURL *string

	if linkData.Title != "" {
		title = &linkData.Title
	}

	if linkData.Description != "" {
		description = &linkData.Description
	}

	if linkData.ImageURL != "" {
		imageURL = &linkData.ImageURL
	}

	// Call AddLinkProperties with the extracted data
	return AddLinkProperties(block, &url, title, description, imageURL, true)
}

func AddLinkProperties(b *Block, url, title, description, urlImage *string, enriched bool) error {
	if b == nil {
		return fmt.Errorf("cannot add link properties because given b is nil")
	}

	// Always set URL property if provided (even if empty)
	if url != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyURL, *url); err != nil {
			return fmt.Errorf("failed to set URL property: %w", err)
		}
	}

	// Always set title property if provided (even if empty)
	if title != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTitle, *title); err != nil {
			return fmt.Errorf("failed to set title property: %w", err)
		}
	}

	// Only set description if not empty
	if description != nil && *description != "" {
		if err := b.Properties.ReplaceValue(PropertyKeyDescription, *description); err != nil {
			return fmt.Errorf("failed to set description property: %w", err)
		}
	}

	// Only set image URL if not empty
	if urlImage != nil && *urlImage != "" {
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
