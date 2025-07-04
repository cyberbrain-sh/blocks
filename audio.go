package blocks

import (
	"fmt"
	"strings"
)

// AddAudioProperties adds audio properties to the given block
func AddAudioProperties(b *Block, size *int, transcription *string, publicURL *string, enriched bool) error {
	if b == nil {
		return fmt.Errorf("cannot add audio properties because given block is nil")
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

// RenderAudioProperties renders audio properties in a human-readable format
func RenderAudioProperties(b Block) string {
	// Get audio-specific properties
	size, hasSize := b.Properties.GetInt(PropertyKeySize)
	transcription, hasTranscription := b.Properties.GetString(PropertyKeyTranscription)
	publicURL, hasPublicURL := b.Properties.GetString(PropertyKeyPublicURL)

	// Build the markdown representation
	var parts []string

	// Add audio header
	parts = append(parts, "## Audio")

	// Add public URL if available
	if hasPublicURL && publicURL != "" {
		parts = append(parts, fmt.Sprintf("[🔊 Audio File](%s)", publicURL))
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

// GetAudioProperties returns a list of all audio property keys
func GetAudioProperties() []string {
	return []string{
		PropertyKeySize,
		PropertyKeyTranscription,
		PropertyKeyPublicURL,
		PropertyKeyEnriched,
	}
}
