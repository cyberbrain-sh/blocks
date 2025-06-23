package blocks

import "testing"

func TestRenderSeriesProperties(t *testing.T) {
	tests := []struct {
		name     string
		block    Block
		expected string
	}{
		{
			name: "basic series with all properties",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:            []interface{}{"Breaking Bad"},
					PropertyKeyURL:              []interface{}{"https://example.com/breakingbad"},
					PropertyKeyDescription:      []interface{}{"A high school chemistry teacher turned methamphetamine producer."},
					PropertyKeyFirstAirYear:     []interface{}{2008},
					PropertyKeyLastAirYear:      []interface{}{2013},
					PropertyKeyRating:           []interface{}{9.5},
					PropertyKeyStatus:           []interface{}{"Ended"},
					PropertyKeyNumberOfSeasons:  []interface{}{5},
					PropertyKeyNumberOfEpisodes: []interface{}{62},
					PropertyKeyGenres:           []interface{}{"Crime", "Drama", "Thriller"},
					PropertyKeyCreators:         []interface{}{"Vince Gilligan"},
					PropertyKeyCast:             []interface{}{"Bryan Cranston", "Aaron Paul"},
					PropertyKeyNetworks:         []interface{}{"AMC"},
					PropertyKeyChecked:          []interface{}{false},
				},
			},
			expected: `- [ ] ## [Breaking Bad (2008-2013)](https://example.com/breakingbad)
**Status:** Ended
**Rating: ⭐ 9.5 | 5 Seasons, 62 Episodes**
**Networks:** AMC
**Genres:** Crime, Drama, Thriller
**Creators:** Vince Gilligan
**Cast:** Bryan Cranston, Aaron Paul

A high school chemistry teacher turned methamphetamine producer.`,
		},
		{
			name: "series with nested arrays",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:        []interface{}{"Stranger Things"},
					PropertyKeyFirstAirYear: []interface{}{2016},
					PropertyKeyGenres:       []interface{}{[]interface{}{"Drama", "Fantasy"}, "Horror"},
					PropertyKeyCreators:     []interface{}{[]string{"The Duffer Brothers"}},
					PropertyKeyCast:         []interface{}{"Winona Ryder", []interface{}{"David Harbour", "Finn Wolfhard"}},
					PropertyKeyNetworks:     []interface{}{[]interface{}{"Netflix"}},
				},
			},
			expected: `## Stranger Things (2016)
**Networks:** Netflix
**Genres:** Drama, Fantasy, Horror
**Creators:** The Duffer Brothers
**Cast:** Winona Ryder, David Harbour, Finn Wolfhard`,
		},
		{
			name: "series with checked status",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:        []interface{}{"The Office"},
					PropertyKeyFirstAirYear: []interface{}{2005},
					PropertyKeyChecked:      []interface{}{true},
				},
			},
			expected: `- [x] ## The Office (2005)`,
		},
		{
			name: "series with string rating and seasons",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:           []interface{}{"Friends"},
					PropertyKeyFirstAirYear:    []interface{}{1994},
					PropertyKeyRating:          []interface{}{"8.9"},
					PropertyKeyNumberOfSeasons: []interface{}{"10"},
				},
			},
			expected: `## Friends (1994)
**Rating: ⭐ 8.9 | 10 Seasons**`,
		},
		{
			name: "series with float64 rating and episodes",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:            []interface{}{"Lost"},
					PropertyKeyFirstAirYear:     []interface{}{2004},
					PropertyKeyRating:           []interface{}{float64(8.3)},
					PropertyKeyNumberOfEpisodes: []interface{}{float64(121)},
				},
			},
			expected: `## Lost (2004)
**Rating: ⭐ 8.3 | 121 Episodes**`,
		},
		{
			name: "series with empty arrays",
			block: Block{
				Properties: Properties{
					PropertyKeyTitle:    []interface{}{"Empty Series"},
					PropertyKeyGenres:   []interface{}{},
					PropertyKeyCreators: []interface{}{},
					PropertyKeyCast:     []interface{}{},
					PropertyKeyNetworks: []interface{}{},
				},
			},
			expected: `## Empty Series`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RenderSeriesProperties(tt.block)
			if result != tt.expected {
				t.Errorf("RenderSeriesProperties() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestFlattenArray_Series(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []string
	}{
		{
			name:     "flat string array",
			input:    []interface{}{"Drama", "Comedy"},
			expected: []string{"Drama", "Comedy"},
		},
		{
			name:     "nested arrays",
			input:    []interface{}{[]interface{}{"Drama", "Fantasy"}, "Horror"},
			expected: []string{"Drama", "Fantasy", "Horror"},
		},
		{
			name:     "mixed types",
			input:    []interface{}{"Drama", []string{"Fantasy", "Thriller"}, 123},
			expected: []string{"Drama", "Fantasy", "Thriller", "123"},
		},
		{
			name:     "empty array",
			input:    []interface{}{},
			expected: []string{},
		},
		{
			name:     "empty strings filtered out",
			input:    []interface{}{"Drama", "", "Comedy", ""},
			expected: []string{"Drama", "Comedy"},
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
