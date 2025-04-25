package pkg

import (
	"fmt"
	"strings"
)

func AddYoutubePropertiesFromPage(block *Block, page PageData) error {
	if block == nil {
		return fmt.Errorf("cannot add youtube properties because given block is nil")
	}

	// Extract basic page information
	title := page.GetTitle()
	description := page.GetDescription()
	imageURL := page.GetImage()
	originalURL := page.GetURL()

	// Extract custom metadata
	customMetadata := page.GetCustomMetadata()

	// Initialize pointers for AddYoutubeProperties
	var videoID, channelID, channelTitle *string
	var publishedAt, viewCount, likeCount, commentCount *string
	var duration, definition *string
	var hasCaptions *bool
	var tags *[]string
	var checked *bool = nil // Add checked property

	if customMetadata != nil {
		// Extract YouTube specific fields
		if val, ok := customMetadata["video_id"]; ok && val != "" {
			videoID = &val
		}

		if val, ok := customMetadata["channel_id"]; ok && val != "" {
			channelID = &val
		}

		if val, ok := customMetadata["channel_title"]; ok && val != "" {
			channelTitle = &val
		}

		if val, ok := customMetadata["published_at"]; ok && val != "" {
			publishedAt = &val
		}

		if val, ok := customMetadata["view_count"]; ok && val != "" {
			viewCount = &val
		}

		if val, ok := customMetadata["like_count"]; ok && val != "" {
			likeCount = &val
		}

		if val, ok := customMetadata["comment_count"]; ok && val != "" {
			commentCount = &val
		}

		if val, ok := customMetadata["duration"]; ok && val != "" {
			duration = &val
		}

		if val, ok := customMetadata["definition"]; ok && val != "" {
			definition = &val
		}

		// Parse has_captions to boolean
		if hasCaptionsStr, ok := customMetadata["has_captions"]; ok {
			hasCaptionsBool := hasCaptionsStr == "true"
			hasCaptions = &hasCaptionsBool
		}

		// Extract tags by splitting on commas
		if tagsStr, ok := customMetadata["tags"]; ok && tagsStr != "" {
			tagsList := strings.Split(tagsStr, ", ")
			tags = &tagsList
		}

		// Parse checked property if available
		if checkedStr, ok := customMetadata["checked"]; ok {
			checkedBool := checkedStr == "true"
			checked = &checkedBool
		}
	}

	// Call AddYoutubeProperties with the extracted values
	return AddYoutubeProperties(block, &title, &description, &imageURL, &originalURL,
		videoID, channelID, channelTitle, publishedAt, viewCount, likeCount, commentCount,
		duration, definition, hasCaptions, tags, checked, true)
}

func AddYoutubeProperties(b *Block, title, description, imageURL, url, videoID, channelID, channelTitle,
	publishedAt, viewCount, likeCount, commentCount, duration, definition *string, hasCaptions *bool, tags *[]string, checked *bool, enriched bool) error {
	if b == nil {
		return fmt.Errorf("cannot add youtube properties because given b is nil")
	}

	// Common properties
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

	if imageURL != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyImageURL, *imageURL); err != nil {
			return fmt.Errorf("failed to set image URL property: %w", err)
		}
	}

	if url != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyURL, *url); err != nil {
			return fmt.Errorf("failed to set URL property: %w", err)
		}
	}

	// YouTube specific properties
	if videoID != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyVideoID, *videoID); err != nil {
			return fmt.Errorf("failed to set video ID property: %w", err)
		}
	}

	if channelID != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyChannelID, *channelID); err != nil {
			return fmt.Errorf("failed to set channel ID property: %w", err)
		}
	}

	if channelTitle != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyChannelTitle, *channelTitle); err != nil {
			return fmt.Errorf("failed to set channel title property: %w", err)
		}
	}

	if publishedAt != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyPublishedAt, *publishedAt); err != nil {
			return fmt.Errorf("failed to set published at property: %w", err)
		}
	}

	if viewCount != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyViewCount, *viewCount); err != nil {
			return fmt.Errorf("failed to set view count property: %w", err)
		}
	}

	if likeCount != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyLikeCount, *likeCount); err != nil {
			return fmt.Errorf("failed to set like count property: %w", err)
		}
	}

	if commentCount != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyCommentCount, *commentCount); err != nil {
			return fmt.Errorf("failed to set comment count property: %w", err)
		}
	}

	if duration != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyDuration, *duration); err != nil {
			return fmt.Errorf("failed to set duration property: %w", err)
		}
	}

	if definition != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyDefinition, *definition); err != nil {
			return fmt.Errorf("failed to set definition property: %w", err)
		}
	}

	if hasCaptions != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyHasCaptions, *hasCaptions); err != nil {
			return fmt.Errorf("failed to set has captions property: %w", err)
		}
	}

	// Checked property
	if checked != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyChecked, *checked); err != nil {
			return fmt.Errorf("failed to set checked property: %w", err)
		}
	}

	if tags != nil && len(*tags) > 0 {
		tagInterfaces := make([]interface{}, len(*tags))
		for i, tag := range *tags {
			tagInterfaces[i] = tag
		}
		if err := b.Properties.ReplaceValue(PropertyKeyTags, tagInterfaces); err != nil {
			return fmt.Errorf("failed to set tags property: %w", err)
		}
	}

	if err := b.Properties.ReplaceValue(PropertyKeyEnriched, enriched); err != nil {
		return fmt.Errorf("failed to set enriched property: %w", err)
	}

	return nil
}

func RenderYoutubeProperties(b Block) string {
	// Get main properties
	titleValue, hasTitle := b.Properties.Get(PropertyKeyTitle)
	title, titleOk := titleValue.(string)

	urlValue, hasURL := b.Properties.Get(PropertyKeyURL)
	url, urlOk := urlValue.(string)

	descValue, hasDesc := b.Properties.Get(PropertyKeyDescription)
	description, descOk := descValue.(string)

	// Get checked property
	checkedValue, hasChecked := b.Properties.Get(PropertyKeyChecked)
	checked, _ := checkedValue.(bool)

	// Get YouTube-specific properties
	channelValue, hasChannel := b.Properties.Get(PropertyKeyChannelTitle)
	channel, channelOk := channelValue.(string)

	viewsValue, hasViews := b.Properties.Get(PropertyKeyViewCount)
	views, viewsOk := viewsValue.(string)

	durationValue, hasDuration := b.Properties.Get(PropertyKeyDuration)
	duration, durationOk := durationValue.(string)

	// Build the markdown representation
	var parts []string

	// Add checkbox if checked property exists
	prefix := ""
	if hasChecked {
		if checked {
			prefix = "- [x] "
		} else {
			prefix = "- [ ] "
		}
	}

	// Title + link
	if hasTitle && titleOk && title != "" {
		if hasURL && urlOk && url != "" {
			parts = append(parts, fmt.Sprintf("%s## [%s](%s)", prefix, title, url))
		} else {
			parts = append(parts, fmt.Sprintf("%s## %s", prefix, title))
		}
	} else if hasURL && urlOk && url != "" {
		parts = append(parts, fmt.Sprintf("%s## [YouTube Video](%s)", prefix, url))
	}

	// Channel info
	if hasChannel && channelOk && channel != "" {
		parts = append(parts, fmt.Sprintf("**Channel:** %s", channel))
	}

	// Stats
	var stats []string
	if hasDuration && durationOk && duration != "" {
		stats = append(stats, fmt.Sprintf("Duration: %s", duration))
	}
	if hasViews && viewsOk && views != "" {
		stats = append(stats, fmt.Sprintf("Views: %s", views))
	}

	if len(stats) > 0 {
		parts = append(parts, fmt.Sprintf("**Stats:** %s", strings.Join(stats, " | ")))
	}

	// Description
	if hasDesc && descOk && description != "" {
		parts = append(parts, fmt.Sprintf("\n%s", description))
	}

	return strings.Join(parts, "\n")
}

func GetYoutubeProperties() []string {
	return []string{
		PropertyKeyTitle,
		PropertyKeyDescription,
		PropertyKeyImageURL,
		PropertyKeyURL,
		PropertyKeyVideoID,
		PropertyKeyChannelID,
		PropertyKeyChannelTitle,
		PropertyKeyPublishedAt,
		PropertyKeyViewCount,
		PropertyKeyLikeCount,
		PropertyKeyCommentCount,
		PropertyKeyDuration,
		PropertyKeyDefinition,
		PropertyKeyHasCaptions,
		PropertyKeyTags,
		PropertyKeyEnriched,
		PropertyKeyChecked,
	}
}
