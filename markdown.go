package blocks

import (
	"fmt"
	"regexp"
	"strings"
)

// NewBlockFromMarkdown creates a new Block from markdown text.
// It determines the block type based on the markdown syntax and sets the appropriate properties.
// Returns the created Block and an error if any.
func NewBlockFromMarkdown(markdown string) (Block, error) {
	// Create a new empty block
	block := NewEmptyBlock()

	// Trim leading/trailing whitespace
	markdown = strings.TrimSpace(markdown)

	// Determine block type and content based on markdown syntax
	if len(markdown) == 0 {
		// Empty markdown, create a paragraph
		block.Type = TypeParagraph
		emptyString := ""
		if err := AddParagraphProperties(&block, &emptyString); err != nil {
			return block, fmt.Errorf("failed to add paragraph properties: %w", err)
		}
		return block, nil
	}

	// Check for markdown links [text](url)
	linkPattern := regexp.MustCompile(`^\[(.*?)\]\((.*?)\)$`)
	if matches := linkPattern.FindStringSubmatch(markdown); len(matches) == 3 {
		block.Type = TypeLink
		text := matches[1]
		url := matches[2]
		enriched := false
		if err := AddLinkProperties(&block, &url, &text, nil, nil, enriched); err != nil {
			return block, fmt.Errorf("failed to add link properties: %w", err)
		}
		return block, nil
	}

	// Check for plain URLs (http://, https://, ftp://)
	urlPattern := regexp.MustCompile(`^(https?|ftp)://\S+$`)
	if urlPattern.MatchString(markdown) {
		block.Type = TypeLink
		url := markdown
		// Use the URL as the title as well
		if err := AddLinkProperties(&block, &url, &url, nil, nil, false); err != nil {
			return block, fmt.Errorf("failed to add link properties: %w", err)
		}
		return block, nil
	}

	// Check for headers (# to ######)
	if strings.HasPrefix(markdown, "# ") {
		block.Type = TypeHeader1
		content := strings.TrimPrefix(markdown, "# ")
		if err := AddHeaderProperties(&block, &content); err != nil {
			return block, fmt.Errorf("failed to add header properties: %w", err)
		}
	} else if strings.HasPrefix(markdown, "## ") {
		block.Type = TypeHeader2
		content := strings.TrimPrefix(markdown, "## ")
		if err := AddHeaderProperties(&block, &content); err != nil {
			return block, fmt.Errorf("failed to add header properties: %w", err)
		}
	} else if strings.HasPrefix(markdown, "### ") {
		block.Type = TypeHeader3
		content := strings.TrimPrefix(markdown, "### ")
		if err := AddHeaderProperties(&block, &content); err != nil {
			return block, fmt.Errorf("failed to add header properties: %w", err)
		}
	} else if strings.HasPrefix(markdown, "#### ") {
		block.Type = TypeHeader4
		content := strings.TrimPrefix(markdown, "#### ")
		if err := AddHeaderProperties(&block, &content); err != nil {
			return block, fmt.Errorf("failed to add header properties: %w", err)
		}
	} else if strings.HasPrefix(markdown, "##### ") {
		block.Type = TypeHeader5
		content := strings.TrimPrefix(markdown, "##### ")
		if err := AddHeaderProperties(&block, &content); err != nil {
			return block, fmt.Errorf("failed to add header properties: %w", err)
		}
	} else if strings.HasPrefix(markdown, "###### ") {
		block.Type = TypeHeader6
		content := strings.TrimPrefix(markdown, "###### ")
		if err := AddHeaderProperties(&block, &content); err != nil {
			return block, fmt.Errorf("failed to add header properties: %w", err)
		}
	} else if strings.HasPrefix(markdown, "- ") || strings.HasPrefix(markdown, "* ") {
		// Bullet list item
		block.Type = TypeBulletListItem
		var content string
		if strings.HasPrefix(markdown, "- ") {
			content = strings.TrimPrefix(markdown, "- ")
		} else {
			content = strings.TrimPrefix(markdown, "* ")
		}
		if err := AddBulletListItemProperties(&block, &content); err != nil {
			return block, fmt.Errorf("failed to add bullet list item properties: %w", err)
		}
	} else if matchNumberedList(markdown) {
		// Numbered list item (e.g., "1. ", "2. ", etc.)
		block.Type = TypeNumberedListItem
		// Extract the content after the number and dot
		parts := strings.SplitN(markdown, ". ", 2)
		if len(parts) == 2 {
			content := parts[1]
			if err := AddNumberedListItemProperties(&block, &content); err != nil {
				return block, fmt.Errorf("failed to add numbered list item properties: %w", err)
			}
		} else {
			// Fallback to empty content if parsing fails
			emptyString := ""
			if err := AddNumberedListItemProperties(&block, &emptyString); err != nil {
				return block, fmt.Errorf("failed to add numbered list item properties: %w", err)
			}
		}
	} else {
		// Default to paragraph
		block.Type = TypeParagraph
		if err := AddParagraphProperties(&block, &markdown); err != nil {
			return block, fmt.Errorf("failed to add paragraph properties: %w", err)
		}
	}

	return block, nil
}

// matchNumberedList checks if a string starts with a number followed by a period and a space
func matchNumberedList(s string) bool {
	for i, c := range s {
		if i == 0 {
			if c < '0' || c > '9' {
				return false
			}
		} else if c == '.' {
			// Check if the next character is a space and we've seen at least one digit
			return i > 0 && i+1 < len(s) && s[i+1] == ' '
		} else if c < '0' || c > '9' {
			return false
		}
	}
	return false
}
