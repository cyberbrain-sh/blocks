package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func AddEmailProperties(block *Block, emailID, threadID, from, to, subject, text *string,
	date, receivedAt *time.Time, attachments, labels *[]string) error {
	if block == nil {
		return fmt.Errorf("cannot add email properties because given block is nil")
	}

	if emailID != nil {
		if err := block.Properties.ReplaceValue(PropertyKeyEmailID, *emailID); err != nil {
			return fmt.Errorf("failed to set email ID property: %w", err)
		}
	}

	if threadID != nil {
		if err := block.Properties.ReplaceValue(PropertyKeyThreadID, *threadID); err != nil {
			return fmt.Errorf("failed to set thread ID property: %w", err)
		}
	}

	if from != nil {
		if err := block.Properties.ReplaceValue(PropertyKeyFrom, *from); err != nil {
			return fmt.Errorf("failed to set from property: %w", err)
		}
	}

	if to != nil {
		if err := block.Properties.ReplaceValue(PropertyKeyTo, *to); err != nil {
			return fmt.Errorf("failed to set to property: %w", err)
		}
	}

	if subject != nil {
		if err := block.Properties.ReplaceValue(PropertyKeySubject, *subject); err != nil {
			return fmt.Errorf("failed to set subject property: %w", err)
		}
	}

	if text != nil {
		if err := block.Properties.ReplaceValue(PropertyKeyText, *text); err != nil {
			return fmt.Errorf("failed to set text property: %w", err)
		}
	}

	if date != nil {
		if err := block.Properties.ReplaceValue(PropertyKeyDate, *date); err != nil {
			return fmt.Errorf("failed to set date property: %w", err)
		}
	}

	if receivedAt != nil {
		if err := block.Properties.ReplaceValue(PropertyKeyReceivedAt, *receivedAt); err != nil {
			return fmt.Errorf("failed to set received at property: %w", err)
		}
	}

	// Array properties
	if attachments != nil && len(*attachments) > 0 {
		attachmentInterfaces := make([]interface{}, len(*attachments))
		for i, attachment := range *attachments {
			attachmentInterfaces[i] = attachment
		}
		if err := block.Properties.ReplaceValue(PropertyKeyAttachments, attachmentInterfaces); err != nil {
			return fmt.Errorf("failed to set attachments property: %w", err)
		}
	}

	if labels != nil && len(*labels) > 0 {
		labelInterfaces := make([]interface{}, len(*labels))
		for i, label := range *labels {
			labelInterfaces[i] = label
		}
		if err := block.Properties.ReplaceValue(PropertyKeyLabels, labelInterfaces); err != nil {
			return fmt.Errorf("failed to set labels property: %w", err)
		}
	}

	return nil
}

// EmailData is an anonymous struct that matches the fields of structured.DataEmail
// This allows us to remove the dependency on the structured package
type EmailData struct {
	ID          string    `json:"id"`
	ThreadID    string    `json:"thread_id"`
	From        string    `json:"from"`
	To          string    `json:"to"`
	Subject     string    `json:"subject"`
	Body        string    `json:"body"`
	Date        time.Time `json:"date"`
	ReceivedAt  time.Time `json:"received_at"`
	Attachments []string  `json:"attachments"`
	Labels      []string  `json:"labels"`
}

func AddEmailPropertiesFromStructured(b *Block, email *EmailData) error {
	if b == nil {
		return fmt.Errorf("cannot add email properties because given b is nil")
	}

	if email == nil {
		return fmt.Errorf("cannot add email properties because given email is nil")
	}

	// Set ID
	if email.ID != "" {
		if err := b.Properties.ReplaceValue(PropertyKeyEmailID, email.ID); err != nil {
			return fmt.Errorf("failed to set email ID property: %w", err)
		}
	}

	// Set ThreadID
	if email.ThreadID != "" {
		if err := b.Properties.ReplaceValue(PropertyKeyThreadID, email.ThreadID); err != nil {
			return fmt.Errorf("failed to set thread ID property: %w", err)
		}
	}

	// Set From
	if email.From != "" {
		if err := b.Properties.ReplaceValue(PropertyKeyFrom, email.From); err != nil {
			return fmt.Errorf("failed to set from property: %w", err)
		}
	}

	// Set To
	if email.To != "" {
		if err := b.Properties.ReplaceValue(PropertyKeyTo, email.To); err != nil {
			return fmt.Errorf("failed to set to property: %w", err)
		}
	}

	// Set Subject
	if email.Subject != "" {
		if err := b.Properties.ReplaceValue(PropertyKeySubject, email.Subject); err != nil {
			return fmt.Errorf("failed to set subject property: %w", err)
		}
	}

	// Set Body/Text
	if email.Body != "" {
		if err := b.Properties.ReplaceValue(PropertyKeyText, email.Body); err != nil {
			return fmt.Errorf("failed to set text property: %w", err)
		}
	}

	// Set Date
	if !email.Date.IsZero() {
		if err := b.Properties.ReplaceValue(PropertyKeyDate, email.Date); err != nil {
			return fmt.Errorf("failed to set date property: %w", err)
		}
	}

	// Set ReceivedAt
	if !email.ReceivedAt.IsZero() {
		if err := b.Properties.ReplaceValue(PropertyKeyReceivedAt, email.ReceivedAt); err != nil {
			return fmt.Errorf("failed to set received at property: %w", err)
		}
	}

	// Set Attachments
	if len(email.Attachments) > 0 {
		attachmentInterfaces := make([]interface{}, len(email.Attachments))
		for i, attachment := range email.Attachments {
			attachmentInterfaces[i] = attachment
		}
		if err := b.Properties.ReplaceValue(PropertyKeyAttachments, attachmentInterfaces); err != nil {
			return fmt.Errorf("failed to set attachments property: %w", err)
		}
	}

	// Set Labels
	if len(email.Labels) > 0 {
		labelInterfaces := make([]interface{}, len(email.Labels))
		for i, label := range email.Labels {
			labelInterfaces[i] = label
		}
		if err := b.Properties.ReplaceValue(PropertyKeyLabels, labelInterfaces); err != nil {
			return fmt.Errorf("failed to set labels property: %w", err)
		}
	}

	return nil
}

func RenderEmailProperties(b Block) string {
	// Get basic email properties
	subjectVal, hasSubject := b.Properties.Get(PropertyKeySubject)
	subject, subjectOk := subjectVal.(string)

	fromVal, hasFrom := b.Properties.Get(PropertyKeyFrom)
	from, fromOk := fromVal.(string)

	toVal, hasTo := b.Properties.Get(PropertyKeyTo)
	to, toOk := toVal.(string)

	textVal, hasText := b.Properties.Get(PropertyKeyText)
	text, textOk := textVal.(string)

	// Get date
	var dateStr string
	dateVal, hasDate := b.Properties.Get(PropertyKeyDate)
	if hasDate {
		switch d := dateVal.(type) {
		case time.Time:
			dateStr = d.Format("Jan 2, 2006 3:04 PM")
		case string:
			dateStr = d
		}
	}

	// Get attachments and labels
	attachmentsArr, hasAttachments := b.Properties.GetArray(PropertyKeyAttachments)
	labelsArr, hasLabels := b.Properties.GetArray(PropertyKeyLabels)

	// Build the markdown representation
	var parts []string

	// Subject as header
	if hasSubject && subjectOk && subject != "" {
		parts = append(parts, fmt.Sprintf("## %s", subject))
	}

	// Metadata
	var metadata []string

	if hasFrom && fromOk && from != "" {
		metadata = append(metadata, fmt.Sprintf("**From:** %s", from))
	}

	if hasTo && toOk && to != "" {
		metadata = append(metadata, fmt.Sprintf("**To:** %s", to))
	}

	if dateStr != "" {
		metadata = append(metadata, fmt.Sprintf("**Date:** %s", dateStr))
	}

	parts = append(parts, strings.Join(metadata, "  \n"))

	// Attachments
	if hasAttachments && len(attachmentsArr) > 0 {
		var attachmentsStrs []string
		for _, a := range attachmentsArr {
			if as, ok := a.(string); ok && as != "" {
				attachmentsStrs = append(attachmentsStrs, as)
			}
		}

		if len(attachmentsStrs) > 0 {
			parts = append(parts, fmt.Sprintf("**Attachments:** %s", strings.Join(attachmentsStrs, ", ")))
		}
	}

	// Labels
	if hasLabels && len(labelsArr) > 0 {
		var labelStrs []string
		for _, l := range labelsArr {
			if ls, ok := l.(string); ok && ls != "" {
				labelStrs = append(labelStrs, fmt.Sprintf("`%s`", ls))
			}
		}

		if len(labelStrs) > 0 {
			parts = append(parts, fmt.Sprintf("**Labels:** %s", strings.Join(labelStrs, " ")))
		}
	}

	// Email body
	if hasText && textOk && text != "" {
		// Add a separator
		parts = append(parts, "---")
		// Add the email text
		parts = append(parts, text)
	}

	return strings.Join(parts, "\n\n")
}

// ParseEmailDataFromJSON parses a json.RawMessage into an EmailData struct
func ParseEmailDataFromJSON(rawJSON json.RawMessage) (*EmailData, error) {
	if len(rawJSON) == 0 {
		return nil, fmt.Errorf("cannot parse email data from empty JSON")
	}

	var emailData EmailData
	if err := json.Unmarshal(rawJSON, &emailData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal email data: %w", err)
	}

	return &emailData, nil
}

// AddEmailPropertiesFromJSON parses a json.RawMessage into an EmailData struct
// and adds the email properties to the given block
func AddEmailPropertiesFromJSON(block *Block, rawJSON json.RawMessage) error {
	if block == nil {
		return fmt.Errorf("cannot add email properties because given block is nil")
	}

	if len(rawJSON) == 0 {
		return fmt.Errorf("cannot add email properties because given JSON is empty")
	}

	emailData, err := ParseEmailDataFromJSON(rawJSON)
	if err != nil {
		return fmt.Errorf("failed to parse email data from JSON: %w", err)
	}

	return AddEmailPropertiesFromStructured(block, emailData)
}

func GetEmailProperties() []string {
	return []string{
		PropertyKeyEmailID,
		PropertyKeyThreadID,
		PropertyKeyFrom,
		PropertyKeyTo,
		PropertyKeySubject,
		PropertyKeyText,
		PropertyKeyDate,
		PropertyKeyReceivedAt,
		PropertyKeyAttachments,
		PropertyKeyLabels,
	}
}
