package blocks

import (
	"testing"
)

func TestRenderMovieProperties(t *testing.T) {
	tests := []struct {
		name     string
		block    Block
		expected string
	}{
		{
			name: "basic movie with all properties",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:       []interface{}{"The Matrix"},
					PropertyKeyURL:         []interface{}{"https://example.com/matrix"},
					PropertyKeyDescription: []interface{}{"A computer hacker learns from mysterious rebels about the true nature of his reality."},
					PropertyKeyReleaseYear: []interface{}{1999},
					PropertyKeyRating:      []interface{}{8.7},
					PropertyKeyRuntime:     []interface{}{136},
					PropertyKeyTagline:     []interface{}{"Welcome to the Real World"},
					PropertyKeyGenres:      []interface{}{"Action", "Sci-Fi"},
					PropertyKeyDirectors:   []interface{}{"Lana Wachowski", "Lilly Wachowski"},
					PropertyKeyCast:        []interface{}{"Keanu Reeves", "Laurence Fishburne"},
					PropertyKeyChecked:     []interface{}{false},
				},
			},
			expected: `- [ ] # [The Matrix (1999)](https://example.com/matrix)
*Welcome to the Real World*
**Rating: ⭐ 8.7 | Runtime: 136 min**
**Genres:** Action, Sci-Fi
**Directors:** Lana Wachowski, Lilly Wachowski
**Cast:** Keanu Reeves, Laurence Fishburne

A computer hacker learns from mysterious rebels about the true nature of his reality.`,
		},
		{
			name: "movie with nested arrays",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:       []interface{}{"Inception"},
					PropertyKeyReleaseYear: []interface{}{2010},
					PropertyKeyRating:      []interface{}{8.8},
					PropertyKeyRuntime:     []interface{}{148},
					PropertyKeyGenres:      []interface{}{[]interface{}{"Action", "Thriller"}, "Sci-Fi"},
					PropertyKeyDirectors:   []interface{}{[]string{"Christopher Nolan"}},
					PropertyKeyCast:        []interface{}{"Leonardo DiCaprio", []interface{}{"Joseph Gordon-Levitt", "Ellen Page"}},
				},
			},
			expected: `# Inception (2010)
**Rating: ⭐ 8.8 | Runtime: 148 min**
**Genres:** Action, Thriller, Sci-Fi
**Directors:** Christopher Nolan
**Cast:** Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page`,
		},
		{
			name: "movie with checked status",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:       []interface{}{"The Shawshank Redemption"},
					PropertyKeyReleaseYear: []interface{}{1994},
					PropertyKeyRating:      []interface{}{9.3},
					PropertyKeyRuntime:     []interface{}{142},
					PropertyKeyChecked:     []interface{}{true},
				},
			},
			expected: `- [x] # The Shawshank Redemption (1994)
**Rating: ⭐ 9.3 | Runtime: 142 min**`,
		},
		{
			name: "movie with string rating and runtime",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:       []interface{}{"Pulp Fiction"},
					PropertyKeyReleaseYear: []interface{}{1994},
					PropertyKeyRating:      []interface{}{"8.9"},
					PropertyKeyRuntime:     []interface{}{"154"},
				},
			},
			expected: `# Pulp Fiction (1994)
**Rating: ⭐ 8.9 | Runtime: 154 min**`,
		},
		{
			name: "movie with float64 rating and runtime",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:       []interface{}{"Fight Club"},
					PropertyKeyReleaseYear: []interface{}{1999},
					PropertyKeyRating:      []interface{}{float64(8.8)},
					PropertyKeyRuntime:     []interface{}{float64(139)},
				},
			},
			expected: `# Fight Club (1999)
**Rating: ⭐ 8.8 | Runtime: 139 min**`,
		},
		{
			name: "movie with empty arrays",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:     []interface{}{"Empty Movie"},
					PropertyKeyGenres:    []interface{}{},
					PropertyKeyDirectors: []interface{}{},
					PropertyKeyCast:      []interface{}{},
				},
			},
			expected: `# Empty Movie`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RenderMovieProperties(tt.block)
			if result != tt.expected {
				t.Errorf("RenderMovieProperties() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestFlattenArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []string
	}{
		{
			name:     "flat string array",
			input:    []interface{}{"Action", "Sci-Fi"},
			expected: []string{"Action", "Sci-Fi"},
		},
		{
			name:     "nested arrays",
			input:    []interface{}{[]interface{}{"Action", "Thriller"}, "Sci-Fi"},
			expected: []string{"Action", "Thriller", "Sci-Fi"},
		},
		{
			name:     "mixed types",
			input:    []interface{}{"Action", []string{"Thriller", "Drama"}, 123},
			expected: []string{"Action", "Thriller", "Drama", "123"},
		},
		{
			name:     "empty array",
			input:    []interface{}{},
			expected: []string{},
		},
		{
			name:     "empty strings filtered out",
			input:    []interface{}{"Action", "", "Sci-Fi", ""},
			expected: []string{"Action", "Sci-Fi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := flattenArray(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("flattenArray() length = %d, want %d", len(result), len(tt.expected))
				return
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("flattenArray()[%d] = %q, want %q", i, v, tt.expected[i])
				}
			}
		})
	}
}
