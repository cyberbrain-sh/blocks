package blocks

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRenderToDoProperties(t *testing.T) {
	// Helper function to create string pointers
	strPtr := func(s string) *string {
		return &s
	}

	// Helper function to create bool pointers
	boolPtr := func(b bool) *bool {
		return &b
	}

	// Helper function to create time.Time pointers
	timePtr := func(t time.Time) *time.Time {
		return &t
	}

	// Helper function to create time.Duration pointers
	durationPtr := func(d time.Duration) *time.Duration {
		return &d
	}

	// Create a fixed test time for consistent testing
	testDate := time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC)
	testOffset := 30 * time.Minute

	t.Run("empty block", func(t *testing.T) {
		// Test with an empty block
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}

		result := RenderToDoProperties(block)
		assert.Equal(t, "", result)
	})

	t.Run("title only", func(t *testing.T) {
		// Create a block with only a title
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := "Buy groceries"
		err := AddToDoProperties(&block, strPtr(title), nil, nil, nil)
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		assert.Equal(t, "- [ ] Buy groceries", result)
	})

	t.Run("title and unchecked", func(t *testing.T) {
		// Create a block with a title and unchecked status
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := "Submit report"
		done := false
		err := AddToDoProperties(&block, strPtr(title), boolPtr(done), nil, nil)
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		assert.Equal(t, "- [ ] Submit report", result)
	})

	t.Run("title and checked", func(t *testing.T) {
		// Create a block with a title and checked status
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := "Call mom"
		done := true
		err := AddToDoProperties(&block, strPtr(title), boolPtr(done), nil, nil)
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		assert.Equal(t, "- [x] Call mom", result)
	})

	t.Run("title and date", func(t *testing.T) {
		// Create a block with a title and date
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := "Project deadline"
		err := AddToDoProperties(&block, strPtr(title), nil, timePtr(testDate), nil)
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		assert.Equal(t, "- [ ] Project deadline (Due: 2023-05-15 14:30)", result)
	})

	t.Run("title and offset", func(t *testing.T) {
		// Create a block with a title and reminder offset
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := "Team meeting"
		err := AddToDoProperties(&block, strPtr(title), nil, nil, durationPtr(testOffset))
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		assert.Equal(t, "- [ ] Team meeting (Reminder: 30m0s before)", result)
	})

	t.Run("title, checked, and date", func(t *testing.T) {
		// Create a block with a title, checked status, and date
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := "Doctor appointment"
		done := false
		err := AddToDoProperties(&block, strPtr(title), boolPtr(done), timePtr(testDate), nil)
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		assert.Equal(t, "- [ ] Doctor appointment (Due: 2023-05-15 14:30)", result)
	})

	t.Run("title, checked, and offset", func(t *testing.T) {
		// Create a block with a title, checked status, and reminder offset
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := "Gym session"
		done := true
		err := AddToDoProperties(&block, strPtr(title), boolPtr(done), nil, durationPtr(testOffset))
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		assert.Equal(t, "- [x] Gym session (Reminder: 30m0s before)", result)
	})

	t.Run("title, date, and offset", func(t *testing.T) {
		// Create a block with a title, date, and reminder offset
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := "Conference call"
		err := AddToDoProperties(&block, strPtr(title), nil, timePtr(testDate), durationPtr(testOffset))
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		expected := "- [ ] Conference call (Due: 2023-05-15 14:30, Reminder: 30m0s before at 2023-05-15 14:00)"
		assert.Equal(t, expected, result)
	})

	t.Run("all properties", func(t *testing.T) {
		// Create a block with all properties
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := "Executive meeting"
		done := false
		err := AddToDoProperties(&block, strPtr(title), boolPtr(done), timePtr(testDate), durationPtr(testOffset))
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		expected := "- [ ] Executive meeting (Due: 2023-05-15 14:30, Reminder: 30m0s before at 2023-05-15 14:00)"
		assert.Equal(t, expected, result)
	})

	t.Run("empty title", func(t *testing.T) {
		// Create a block with an empty title
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := ""
		err := AddToDoProperties(&block, strPtr(title), nil, nil, nil)
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		assert.Equal(t, "", result)
	})

	t.Run("different time zones", func(t *testing.T) {
		// Create a block with a date in a different timezone
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := "International call"
		// Create a time with a specific timezone
		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			t.Skip("Skipping timezone test due to location load error:", err)
			return
		}
		nyTime := time.Date(2023, 5, 15, 10, 30, 0, 0, loc)
		err = AddToDoProperties(&block, strPtr(title), nil, timePtr(nyTime), durationPtr(testOffset))
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		// The formatting should preserve the time as specified, though displayed in local format
		assert.Contains(t, result, "- [ ] International call (Due: 2023-05-15 10:30")
		assert.Contains(t, result, "Reminder: 30m0s before at 2023-05-15 10:00")
	})

	t.Run("long offset duration", func(t *testing.T) {
		// Create a block with a longer reminder offset
		block := Block{
			ID:         uuid.New(),
			Type:       TypeToDo,
			Properties: Properties{},
		}
		title := "Annual review"
		longOffset := 24 * time.Hour // 1 day
		err := AddToDoProperties(&block, strPtr(title), nil, timePtr(testDate), durationPtr(longOffset))
		assert.NoError(t, err)

		result := RenderToDoProperties(block)
		expected := "- [ ] Annual review (Due: 2023-05-15 14:30, Reminder: 24h0m0s before at 2023-05-14 14:30)"
		assert.Equal(t, expected, result)
	})
}
