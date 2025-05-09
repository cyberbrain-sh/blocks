package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlockFromMarkdown(t *testing.T) {
	tests := []struct {
		name           string
		markdown       string
		expectedType   DataType
		expectedTitle  string
		expectError    bool
		errorSubstring string
	}{
		{
			name:          "empty string creates paragraph",
			markdown:      "",
			expectedType:  TypeParagraph,
			expectedTitle: "",
			expectError:   false,
		},
		{
			name:          "whitespace only creates paragraph",
			markdown:      "   \t  \n  ",
			expectedType:  TypeParagraph,
			expectedTitle: "",
			expectError:   false,
		},
		{
			name:          "plain text creates paragraph",
			markdown:      "This is a simple paragraph",
			expectedType:  TypeParagraph,
			expectedTitle: "This is a simple paragraph",
			expectError:   false,
		},
		{
			name:          "h1 header",
			markdown:      "# Header 1",
			expectedType:  TypeHeader1,
			expectedTitle: "Header 1",
			expectError:   false,
		},
		{
			name:          "h2 header",
			markdown:      "## Header 2",
			expectedType:  TypeHeader2,
			expectedTitle: "Header 2",
			expectError:   false,
		},
		{
			name:          "h3 header",
			markdown:      "### Header 3",
			expectedType:  TypeHeader3,
			expectedTitle: "Header 3",
			expectError:   false,
		},
		{
			name:          "h4 header",
			markdown:      "#### Header 4",
			expectedType:  TypeHeader4,
			expectedTitle: "Header 4",
			expectError:   false,
		},
		{
			name:          "h5 header",
			markdown:      "##### Header 5",
			expectedType:  TypeHeader5,
			expectedTitle: "Header 5",
			expectError:   false,
		},
		{
			name:          "h6 header",
			markdown:      "###### Header 6",
			expectedType:  TypeHeader6,
			expectedTitle: "Header 6",
			expectError:   false,
		},
		{
			name:          "bullet list with dash",
			markdown:      "- List item with dash",
			expectedType:  TypeBulletListItem,
			expectedTitle: "List item with dash",
			expectError:   false,
		},
		{
			name:          "bullet list with asterisk",
			markdown:      "* List item with asterisk",
			expectedType:  TypeBulletListItem,
			expectedTitle: "List item with asterisk",
			expectError:   false,
		},
		{
			name:          "numbered list",
			markdown:      "1. Numbered list item",
			expectedType:  TypeNumberedListItem,
			expectedTitle: "Numbered list item",
			expectError:   false,
		},
		{
			name:          "numbered list with multiple digits",
			markdown:      "42. Numbered list item with larger number",
			expectedType:  TypeNumberedListItem,
			expectedTitle: "Numbered list item with larger number",
			expectError:   false,
		},
		{
			name:          "hash without space is paragraph",
			markdown:      "#Not a header",
			expectedType:  TypeParagraph,
			expectedTitle: "#Not a header",
			expectError:   false,
		},
		{
			name:          "dash without space is paragraph",
			markdown:      "-Not a list item",
			expectedType:  TypeParagraph,
			expectedTitle: "-Not a list item",
			expectError:   false,
		},
		{
			name:          "number without dot and space is paragraph",
			markdown:      "1Not a numbered list",
			expectedType:  TypeParagraph,
			expectedTitle: "1Not a numbered list",
			expectError:   false,
		},
		{
			name:          "number with dot but no space is paragraph",
			markdown:      "1.Not a numbered list",
			expectedType:  TypeParagraph,
			expectedTitle: "1.Not a numbered list",
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, err := NewBlockFromMarkdown(tt.markdown)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorSubstring != "" {
					assert.Contains(t, err.Error(), tt.errorSubstring)
				}
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedType, block.Type)

			// Check title property based on block type
			var title string
			var ok bool

			title, ok = block.Properties.GetString(PropertyKeyTitle)
			assert.True(t, ok, "Title property should exist")
			assert.Equal(t, tt.expectedTitle, title)
		})
	}
}

// Test edge cases and error handling
func TestNewBlockFromMarkdownEdgeCases(t *testing.T) {
	// Test with a very long markdown string
	longText := "This is a very long paragraph " + string(make([]byte, 10000)) + " that continues."
	block, err := NewBlockFromMarkdown(longText)
	assert.NoError(t, err)
	assert.Equal(t, TypeParagraph, block.Type)

	// Test with special characters
	specialChars := "Special characters: !@#$%^&*()_+{}|:<>?~`-=[]\\;',./😀🚀"
	block, err = NewBlockFromMarkdown(specialChars)
	assert.NoError(t, err)
	assert.Equal(t, TypeParagraph, block.Type)

	// Test with multiple lines
	multiline := "First line\nSecond line\nThird line"
	block, err = NewBlockFromMarkdown(multiline)
	assert.NoError(t, err)
	assert.Equal(t, TypeParagraph, block.Type)
}

// Test the helper function matchNumberedList
func TestMatchNumberedList(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"1. Item", true},
		{"42. Item", true},
		{"0. Item", true},
		{"1.Item", false},   // No space after dot
		{"a. Item", false},  // Doesn't start with number
		{"1", false},        // No dot
		{"1.", false},       // No space after dot
		{". Item", false},   // No number before dot
		{"1 . Item", false}, // Space before dot
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := matchNumberedList(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
