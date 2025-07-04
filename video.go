package blocks

import (
	"fmt"
	"strings"
)

// AddVideoProperties adds video properties to the given block
func AddVideoProperties(b *Block, size *int, transcription *string, publicURL *string, filename *string, extension *string, enriched bool) error {
	if b == nil {
		return fmt.Errorf("cannot add video properties because given block is nil")
	}

	// Set size if provided
	if size != nil {
		if err := b.Properties.ReplaceValue(PropertyKeySize, *size); err != nil {
			return fmt.Errorf("failed to set size property: %w", err)
		}
	}

	// Set transcription if provided
	if transcription != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTranscription, *transcription); err != nil {
			return fmt.Errorf("failed to set transcription property: %w", err)
		}
	}

	// Set public URL if provided
	if publicURL != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyPublicURL, *publicURL); err != nil {
			return fmt.Errorf("failed to set public URL property: %w", err)
		}
	}

	// Set filename if provided
	if filename != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyFilename, *filename); err != nil {
			return fmt.Errorf("failed to set filename property: %w", err)
		}
	}

	// Set extension if provided
	if extension != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyExtension, *extension); err != nil {
			return fmt.Errorf("failed to set extension property: %w", err)
		}
	}

	// Set enriched status
	if err := b.Properties.ReplaceValue(PropertyKeyEnriched, enriched); err != nil {
		return fmt.Errorf("failed to set enriched property: %w", err)
	}

	return nil
}

// RenderVideoProperties renders video properties in a human-readable format
func RenderVideoProperties(b Block) string {
	// Get video-specific properties
	size, hasSize := b.Properties.GetInt(PropertyKeySize)
	transcription, hasTranscription := b.Properties.GetString(PropertyKeyTranscription)
	publicURL, hasPublicURL := b.Properties.GetString(PropertyKeyPublicURL)
	filename, hasFilename := b.Properties.GetString(PropertyKeyFilename)
	extension, hasExtension := b.Properties.GetString(PropertyKeyExtension)

	// Build the markdown representation
	var parts []string

	// Add video header
	parts = append(parts, "## Video")

	// Add filename and extension if available
	if hasFilename && filename != "" {
		headerText := filename
		if hasExtension && extension != "" {
			headerText = fmt.Sprintf("%s.%s", filename, extension)
		}
		parts = append(parts, fmt.Sprintf("**File:** %s", headerText))
	}

	// Add public URL if available
	if hasPublicURL && publicURL != "" {
		parts = append(parts, fmt.Sprintf("[🎥 Video File](%s)", publicURL))
	}

	// Add size information if available
	if hasSize && size > 0 {
		parts = append(parts, fmt.Sprintf("**Size:** %d bytes", size))
	}

	// Add transcription if available
	if hasTranscription && transcription != "" {
		parts = append(parts, fmt.Sprintf("**Transcription:** %s", transcription))
	}

	return strings.Join(parts, "\n")
}

// GetVideoProperties returns a list of all video property keys
func GetVideoProperties() []string {
	return []string{
		PropertyKeySize,
		PropertyKeyTranscription,
		PropertyKeyPublicURL,
		PropertyKeyFilename,
		PropertyKeyExtension,
		PropertyKeyEnriched,
	}
}
