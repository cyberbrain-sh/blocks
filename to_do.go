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

	if done != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyChecked, *done); err != nil {
			return fmt.Errorf("failed to set checked property: %w", err)
		}
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

	// Get reminder offset
	offsetValue, hasOffset := b.Properties.Get(PropertyKeyReminderOffset)
	offset, offsetOk := offsetValue.(time.Duration)

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
			if hasOffset && offsetOk {
				// Calculate the reminder time
				reminderTime := date.Add(-offset)
				result += fmt.Sprintf(", Reminder: %s before at %s",
					offset,
					reminderTime.Format("2006-01-02 15:04"))
			}

			result += ")"
		} else if hasOffset && offsetOk {
			// Only offset without date (can't calculate reminder time)
			result += fmt.Sprintf(" (Reminder: %s before)", offset)
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
