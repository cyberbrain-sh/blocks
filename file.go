package blocks

import (
	"fmt"
	"strings"
)

// AddFileProperties adds file properties to the given block
func AddFileProperties(b *Block, size *int, transcription *string, publicURL *string, enriched bool) error {
	if b == nil {
		return fmt.Errorf("cannot add file properties because given block is nil")
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

	// Set enriched status
	if err := b.Properties.ReplaceValue(PropertyKeyEnriched, enriched); err != nil {
		return fmt.Errorf("failed to set enriched property: %w", err)
	}

	return nil
}

// RenderFileProperties renders file properties in a human-readable format
func RenderFileProperties(b Block) string {
	// Get file-specific properties
	size, hasSize := b.Properties.GetInt(PropertyKeySize)
	transcription, hasTranscription := b.Properties.GetString(PropertyKeyTranscription)
	publicURL, hasPublicURL := b.Properties.GetString(PropertyKeyPublicURL)

	// Build the markdown representation
	var parts []string

	// Add file header
	parts = append(parts, "## File")

	// Add public URL if available
	if hasPublicURL && publicURL != "" {
		parts = append(parts, fmt.Sprintf("[📄 File](%s)", publicURL))
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

// GetFileProperties returns a list of all file property keys
func GetFileProperties() []string {
	return []string{
		PropertyKeySize,
		PropertyKeyTranscription,
		PropertyKeyPublicURL,
		PropertyKeyEnriched,
	}
}
