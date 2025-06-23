package blocks

import (
	"fmt"
	"time"
)

func AddToDoProperties(b *Block, title *string, done *bool, date *time.Time, offset *time.Duration) error {
	if b == nil {
		return fmt.Errorf("cannot add task properties because given b is nil")
	}

	if title != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTitle, *title); err != nil {
			return fmt.Errorf("failed to set title property: %w", err)
		}
	}

	// Always set checked property, defaulting to false if done is nil
	if done == nil {
		pFalse := false
		done = &pFalse
	}
	if err := b.Properties.ReplaceValue(PropertyKeyChecked, *done); err != nil {
		return fmt.Errorf("failed to set checked property: %w", err)
	}

	if date != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTargetDateTime, *date); err != nil {
			return fmt.Errorf("failed to set target date/time property: %w", err)
		}
	}

	if offset != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyReminderOffset, *offset); err != nil {
			return fmt.Errorf("failed to set reminder offset property: %w", err)
		}
	}

	return nil
}

func RenderToDoProperties(b Block) string {
	// Get task title
	titleValue, hasTitle := b.Properties.Get(PropertyKeyTitle)
	title, titleOk := titleValue.(string)

	// Get checked status
	checkedValue, hasChecked := b.Properties.Get(PropertyKeyChecked)
	checked, checkedOk := checkedValue.(bool)

	// Get target date/time
	dateValue, hasDate := b.Properties.Get(PropertyKeyTargetDateTime)
	date, dateOk := dateValue.(time.Time)

	// Get reminder offset (support int, float64, or time.Duration)
	offsetValue, hasOffset := b.Properties.Get(PropertyKeyReminderOffset)
	var offset time.Duration
	var offsetOk bool
	if hasOffset {
		switch v := offsetValue.(type) {
		case time.Duration:
			offset = v
			offsetOk = true
		case int:
			offset = time.Duration(v) * time.Second
			offsetOk = true
		case float64:
			offset = time.Duration(int(v)) * time.Second
			offsetOk = true
		case string:
			// Ignore string parsing for now, as offset should be int, float64, or time.Duration
			// Could add parsing logic if needed
			// No-op
		}
	}

	// Format as markdown task
	if hasTitle && titleOk && title != "" {
		var result string

		// Use markdown checkbox syntax for the base task
		if hasChecked && checkedOk {
			if checked {
				result = fmt.Sprintf("- [x] %s", title) // Checked task
			} else {
				result = fmt.Sprintf("- [ ] %s", title) // Unchecked task
			}
		} else {
			// No check status, just render the title
			result = fmt.Sprintf("- %s", title)
		}

		// Add date if available
		if hasDate && dateOk {
			result += fmt.Sprintf(" (Due: %s", date.Format("2006-01-02 15:04"))

			// Add offset and reminder time if available
			if hasOffset && offsetOk && offset != 0 {
				reminderTime := date.Add(-offset)
				result += fmt.Sprintf(", Reminder: %s before at %s",
					offset.String(),
					reminderTime.Format("2006-01-02 15:04"))
			}

			result += ")"
		} else if hasOffset && offsetOk && offset != 0 {
			// Only offset without date (can't calculate reminder time)
			result += fmt.Sprintf(" (Reminder: %s before)", offset.String())
		}

		return result
	}

	return ""
}

func GetToDoProperties() []string {
	return []string{
		PropertyKeyTitle,
		PropertyKeyChecked,
		PropertyKeyTargetDateTime,
		PropertyKeyReminderOffset,
	}
}
