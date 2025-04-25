package pkg

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProperties_MarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		properties Properties
		want       string
		wantErr    bool
	}{
		{
			name:       "empty properties",
			properties: Properties{},
			want:       "{}",
			wantErr:    false,
		},
		{
			name: "properties with simple values",
			properties: Properties{
				PropertyKeyTitle: {"Test Title"},
				PropertyKeyText:  {"Some text content"},
			},
			want:    `{"text":["Some text content"],"title":["Test Title"]}`,
			wantErr: false,
		},
		{
			name: "properties with mixed value types",
			properties: Properties{
				PropertyKeyTitle:     {"Title"},
				PropertyKeyChecked:   {true},
				PropertyKeyViewCount: {1000},
			},
			want:    `{"checked":[true],"title":["Title"],"view_count":[1000]}`,
			wantErr: false,
		},
		{
			name: "properties with multiple values",
			properties: Properties{
				PropertyKeyGenres:  {"Action", "Adventure", "Sci-Fi"},
				PropertyKeyTagline: {"In space no one can hear you scream"},
			},
			want:    `{"genres":["Action","Adventure","Sci-Fi"],"tagline":["In space no one can hear you scream"]}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.properties)

			if (err != nil) != tt.wantErr {
				t.Errorf("Properties.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Compare JSON strings allowing for different order of keys
			var gotMap, wantMap map[string]interface{}
			if err := json.Unmarshal(got, &gotMap); err != nil {
				t.Errorf("couldn't unmarshal result: %v", err)
				return
			}
			if err := json.Unmarshal([]byte(tt.want), &wantMap); err != nil {
				t.Errorf("couldn't unmarshal expected: %v", err)
				return
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("Properties.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestProperties_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		jsonString string
		want       Properties
		wantErr    bool
	}{
		{
			name:       "empty JSON object",
			jsonString: "{}",
			want:       Properties{},
			wantErr:    false,
		},
		{
			name:       "simple properties",
			jsonString: `{"title":["Movie Title"],"description":["Movie description"]}`,
			want: Properties{
				PropertyKeyTitle:       {"Movie Title"},
				PropertyKeyDescription: {"Movie description"},
			},
			wantErr: false,
		},
		{
			name:       "properties with mixed types",
			jsonString: `{"title":["The Matrix"],"release_year":[1999],"rating":[8.7],"in_production":[false]}`,
			want: Properties{
				PropertyKeyTitle:        {"The Matrix"},
				PropertyKeyReleaseYear:  {float64(1999)}, // JSON numbers are parsed as float64 by default
				PropertyKeyRating:       {float64(8.7)},
				PropertyKeyInProduction: {false},
			},
			wantErr: false,
		},
		{
			name:       "properties with array values",
			jsonString: `{"directors":["Lana Wachowski","Lilly Wachowski"],"genres":["Action","Sci-Fi"]}`,
			want: Properties{
				PropertyKeyDirectors: {"Lana Wachowski", "Lilly Wachowski"},
				PropertyKeyGenres:    {"Action", "Sci-Fi"},
			},
			wantErr: false,
		},
		{
			name:       "invalid JSON",
			jsonString: `{"title":"Missing array brackets"}`,
			want:       nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p Properties
			err := json.Unmarshal([]byte(tt.jsonString), &p)

			if (err != nil) != tt.wantErr {
				t.Errorf("Properties.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Properties.UnmarshalJSON() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestProperties_RoundTrip(t *testing.T) {
	// Test that marshaling and then unmarshaling preserves the original data
	original := Properties{
		PropertyKeyTitle:     {"Inception"},
		PropertyKeyDirectors: {"Christopher Nolan"},
		// Use float64 instead of int for numeric values since JSON unmarshals to float64
		PropertyKeyReleaseYear: {float64(2010)},
		PropertyKeyGenres:      {"Action", "Adventure", "Sci-Fi"},
		PropertyKeyRating:      {8.8},
		PropertyKeyRuntime:     {float64(148)},
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	// Unmarshal back to Properties
	var result Properties
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Convert to JSON and back for proper comparison
	originalJSON, _ := json.Marshal(original)
	resultJSON, _ := json.Marshal(result)

	var originalMap, resultMap map[string]interface{}
	json.Unmarshal(originalJSON, &originalMap)
	json.Unmarshal(resultJSON, &resultMap)

	if !reflect.DeepEqual(originalMap, resultMap) {
		t.Errorf("Round-trip failed. Original: %v, Result: %v", original, result)
	}
}

func TestUpdateFromJSON(t *testing.T) {
	t.Run("update multiple fields", func(t *testing.T) {
		// Create a block with initial values
		id := uuid.New()
		spaceID := uuid.New()
		accountID := uuid.New()
		creatorUserID := uuid.New()
		originalSpaceID := uuid.New()
		previousSpaceID := uuid.New()

		createdAt := time.Now().Add(-24 * time.Hour)
		updatedAt := time.Now().Add(-2 * time.Hour)

		block := &Block{
			ID:              id,
			Type:            TypeFragment,
			OriginalSpaceID: originalSpaceID,
			ParentID:        nil,
			AccountID:       accountID,
			SpaceID:         spaceID,
			PreviousSpaceID: previousSpaceID,
			CreatorUserID:   creatorUserID,
			Properties:      Properties{"title": {"Initial Title"}},
			Content:         []uuid.UUID{},
			RawBody:         "Initial body",
			LifecycleStatus: LifecycleStatusCreated,
			Origin:          NewOriginWebapp(),
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}

		// Create update JSON with new values
		newParentID := uuid.New()
		updateJSON := []byte(`{
			"parent_id": "` + newParentID.String() + `",
			"raw_body": "Updated body content",
			"properties": {"title": ["Updated Title"], "tags": ["tag1", "tag2"]}
		}`)

		// Apply update
		updatedFields, err := block.UpdateFromJSON(updateJSON)

		// Assertions
		assert.NoError(t, err)
		assert.Contains(t, updatedFields, "parent_id")
		assert.Contains(t, updatedFields, "raw_body")
		assert.Contains(t, updatedFields, "properties")
		assert.Contains(t, updatedFields, "updated_at")

		assert.Equal(t, &newParentID, block.ParentID)
		assert.Equal(t, "Updated body content", block.RawBody)
		assert.Equal(t, Properties{
			"title": {"Updated Title"},
			"tags":  {"tag1", "tag2"},
		}, block.Properties)
		assert.NotEqual(t, updatedAt, block.UpdatedAt)
	})

	t.Run("update with empty JSON", func(t *testing.T) {
		block := &Block{
			ID:        uuid.New(),
			Type:      TypeFragment,
			RawBody:   "Original body",
			UpdatedAt: time.Now().Add(-1 * time.Hour),
		}

		originalUpdatedAt := block.UpdatedAt

		updatedFields, err := block.UpdateFromJSON([]byte(`{}`))

		assert.NoError(t, err)
		assert.Empty(t, updatedFields)
		assert.Equal(t, originalUpdatedAt, block.UpdatedAt)
	})

	t.Run("update with invalid JSON", func(t *testing.T) {
		block := &Block{
			ID:        uuid.New(),
			Type:      TypeFragment,
			RawBody:   "Original body",
			UpdatedAt: time.Now().Add(-1 * time.Hour),
		}

		updatedFields, err := block.UpdateFromJSON([]byte(`{invalid json`))

		assert.Error(t, err)
		assert.Nil(t, updatedFields)
	})

	t.Run("update with same values (no changes)", func(t *testing.T) {
		spaceID := uuid.New()
		rawBody := "Original body"

		block := &Block{
			ID:        uuid.New(),
			Type:      TypeFragment,
			SpaceID:   spaceID,
			RawBody:   rawBody,
			UpdatedAt: time.Now().Add(-1 * time.Hour),
		}

		originalUpdatedAt := block.UpdatedAt

		updatedFields, err := block.UpdateFromJSON([]byte(`{
			"space_id": "` + spaceID.String() + `",
			"raw_body": "` + rawBody + `"
		}`))

		assert.NoError(t, err)
		assert.Empty(t, updatedFields)
		assert.Equal(t, originalUpdatedAt, block.UpdatedAt)
	})

	t.Run("update array fields", func(t *testing.T) {
		id1 := uuid.New()
		id2 := uuid.New()
		id3 := uuid.New()

		block := &Block{
			ID:      uuid.New(),
			Type:    TypeFragment,
			Content: []uuid.UUID{id1, id2},
		}

		updatedFields, err := block.UpdateFromJSON([]byte(`{
			"content": ["` + id1.String() + `", "` + id3.String() + `"]
		}`))

		assert.NoError(t, err)
		assert.Contains(t, updatedFields, "content")
		assert.Contains(t, updatedFields, "updated_at")
		assert.Equal(t, []uuid.UUID{id1, id3}, block.Content)
	})

	t.Run("update omitempty fields - setting value", func(t *testing.T) {
		block := &Block{
			ID:           uuid.New(),
			Type:         TypeFragment,
			ParentID:     nil,
			LastViewedAt: nil,
		}

		parentID := uuid.New()
		viewedAt := time.Now().Round(time.Second) // Round to avoid fractional second comparison issues

		updatedFields, err := block.UpdateFromJSON([]byte(`{
			"parent_id": "` + parentID.String() + `",
			"last_viewed_at": "` + viewedAt.Format(time.RFC3339) + `"
		}`))

		assert.NoError(t, err)
		assert.Contains(t, updatedFields, "parent_id")
		assert.Contains(t, updatedFields, "last_viewed_at")
		assert.Equal(t, &parentID, block.ParentID)
		assert.Equal(t, viewedAt.UTC(), block.LastViewedAt.UTC())
	})

	t.Run("update omitempty fields - removing value", func(t *testing.T) {
		parentID := uuid.New()
		viewedAt := time.Now()

		block := &Block{
			ID:           uuid.New(),
			Type:         TypeFragment,
			ParentID:     &parentID,
			LastViewedAt: &viewedAt,
		}

		// Note: In typical JSON unmarshaling, omitted fields (not null fields) are not
		// modified. Our function should skip any zero-value fields to emulate this behavior.
		// This test verifies the fields aren't changed when the fields aren't present in JSON.
		updatedFields, err := block.UpdateFromJSON([]byte(`{
			"raw_body": "New content"
		}`))

		assert.NoError(t, err)
		assert.NotContains(t, updatedFields, "parent_id")
		assert.NotContains(t, updatedFields, "last_viewed_at")
		assert.Equal(t, &parentID, block.ParentID)
		assert.Equal(t, &viewedAt, block.LastViewedAt)
	})
}

func TestGetJSONFieldName(t *testing.T) {
	type TestStruct struct {
		RegularField   string `json:"regular_field"`
		OmitemptyField string `json:"omitempty_field,omitempty"`
		DashField      string `json:"-"`
		NoTagField     string
		EmptyTagField  string `json:""`
		ComplicatedTag string `json:"complicated,omitempty,string"`
	}

	testType := reflect.TypeOf(TestStruct{})

	testCases := []struct {
		fieldName string
		expected  string
	}{
		{"RegularField", "regular_field"},
		{"OmitemptyField", "omitempty_field"},
		{"DashField", "DashField"}, // Dash means exclude from JSON, so we fall back to struct field name
		{"NoTagField", "NoTagField"},
		{"EmptyTagField", "EmptyTagField"},
		{"ComplicatedTag", "complicated"},
	}

	for _, tc := range testCases {
		t.Run(tc.fieldName, func(t *testing.T) {
			field, _ := testType.FieldByName(tc.fieldName)
			result := getJSONFieldName(field)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestBlock_RemoveChild(t *testing.T) {
	tests := []struct {
		name            string
		initialBlock    func() Block
		childIDToRemove uuid.UUID
		expectError     bool
		expectedContent func(block Block) []uuid.UUID
	}{
		{
			name: "remove existing child from middle",
			initialBlock: func() Block {
				child1 := uuid.MustParse("11111111-1111-1111-1111-111111111111")
				child2 := uuid.MustParse("22222222-2222-2222-2222-222222222222")
				child3 := uuid.MustParse("33333333-3333-3333-3333-333333333333")

				b := NewEmptyBlock()
				b.Content = []uuid.UUID{child1, child2, child3}
				return b
			},
			childIDToRemove: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			expectError:     false,
			expectedContent: func(block Block) []uuid.UUID {
				return []uuid.UUID{
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uuid.MustParse("33333333-3333-3333-3333-333333333333"),
				}
			},
		},
		{
			name: "remove existing child from beginning",
			initialBlock: func() Block {
				child1 := uuid.MustParse("11111111-1111-1111-1111-111111111111")
				child2 := uuid.MustParse("22222222-2222-2222-2222-222222222222")

				b := NewEmptyBlock()
				b.Content = []uuid.UUID{child1, child2}
				return b
			},
			childIDToRemove: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			expectError:     false,
			expectedContent: func(block Block) []uuid.UUID {
				return []uuid.UUID{
					uuid.MustParse("22222222-2222-2222-2222-222222222222"),
				}
			},
		},
		{
			name: "remove existing child from end",
			initialBlock: func() Block {
				child1 := uuid.MustParse("11111111-1111-1111-1111-111111111111")
				child2 := uuid.MustParse("22222222-2222-2222-2222-222222222222")

				b := NewEmptyBlock()
				b.Content = []uuid.UUID{child1, child2}
				return b
			},
			childIDToRemove: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			expectError:     false,
			expectedContent: func(block Block) []uuid.UUID {
				return []uuid.UUID{
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				}
			},
		},
		{
			name: "remove from empty content",
			initialBlock: func() Block {
				b := NewEmptyBlock()
				b.Content = []uuid.UUID{}
				return b
			},
			childIDToRemove: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			expectError:     true,
			expectedContent: func(block Block) []uuid.UUID {
				return []uuid.UUID{}
			},
		},
		{
			name: "remove non-existent child",
			initialBlock: func() Block {
				child1 := uuid.MustParse("11111111-1111-1111-1111-111111111111")

				b := NewEmptyBlock()
				b.Content = []uuid.UUID{child1}
				return b
			},
			childIDToRemove: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			expectError:     true,
			expectedContent: func(block Block) []uuid.UUID {
				return []uuid.UUID{
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			block := tc.initialBlock()
			initialLength := len(block.Content)

			err := block.RemoveChild(tc.childIDToRemove)

			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, initialLength, len(block.Content), "Content length should not change when an error occurs")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, initialLength-1, len(block.Content), "Content length should decrease by 1 on successful removal")
			}

			assert.Equal(t, tc.expectedContent(block), block.Content)
		})
	}
}

func TestBlock_InsertChild(t *testing.T) {
	tests := []struct {
		name            string
		initialBlock    func() Block
		childIDToInsert uuid.UUID
		afterID         uuid.UUID
		expectError     bool
		expectedContent func(block Block) []uuid.UUID
	}{
		{
			name: "insert after first element",
			initialBlock: func() Block {
				child1 := uuid.MustParse("11111111-1111-1111-1111-111111111111")
				child2 := uuid.MustParse("22222222-2222-2222-2222-222222222222")

				b := NewEmptyBlock()
				b.Content = []uuid.UUID{child1, child2}
				return b
			},
			childIDToInsert: uuid.MustParse("33333333-3333-3333-3333-333333333333"),
			afterID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			expectError:     false,
			expectedContent: func(block Block) []uuid.UUID {
				return []uuid.UUID{
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uuid.MustParse("33333333-3333-3333-3333-333333333333"),
					uuid.MustParse("22222222-2222-2222-2222-222222222222"),
				}
			},
		},
		{
			name: "insert after last element",
			initialBlock: func() Block {
				child1 := uuid.MustParse("11111111-1111-1111-1111-111111111111")
				child2 := uuid.MustParse("22222222-2222-2222-2222-222222222222")

				b := NewEmptyBlock()
				b.Content = []uuid.UUID{child1, child2}
				return b
			},
			childIDToInsert: uuid.MustParse("33333333-3333-3333-3333-333333333333"),
			afterID:         uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			expectError:     false,
			expectedContent: func(block Block) []uuid.UUID {
				return []uuid.UUID{
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uuid.MustParse("22222222-2222-2222-2222-222222222222"),
					uuid.MustParse("33333333-3333-3333-3333-333333333333"),
				}
			},
		},
		{
			name: "insert into empty content",
			initialBlock: func() Block {
				b := NewEmptyBlock()
				b.Content = []uuid.UUID{}
				return b
			},
			childIDToInsert: uuid.MustParse("33333333-3333-3333-3333-333333333333"),
			afterID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			expectError:     true,
			expectedContent: func(block Block) []uuid.UUID {
				return []uuid.UUID{}
			},
		},
		{
			name: "insert after non-existent element",
			initialBlock: func() Block {
				child1 := uuid.MustParse("11111111-1111-1111-1111-111111111111")

				b := NewEmptyBlock()
				b.Content = []uuid.UUID{child1}
				return b
			},
			childIDToInsert: uuid.MustParse("33333333-3333-3333-3333-333333333333"),
			afterID:         uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			expectError:     true,
			expectedContent: func(block Block) []uuid.UUID {
				return []uuid.UUID{
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			block := tc.initialBlock()
			initialLength := len(block.Content)

			err := block.InsertChild(tc.childIDToInsert, tc.afterID)

			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, initialLength, len(block.Content), "Content length should not change when an error occurs")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, initialLength+1, len(block.Content), "Content length should increase by 1 on successful insertion")
			}

			assert.Equal(t, tc.expectedContent(block), block.Content)
		})
	}
}
