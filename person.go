package blocks

import (
	"fmt"
	"strings"
	"time"
)

// AddPersonProperties adds person properties to the given block
func AddPersonProperties(b *Block, firstName *string, lastName *string, birthday *time.Time,
	relationType *string, address *string, phoneNumber *string, imageURL *string, description *string) error {

	if b == nil {
		return fmt.Errorf("cannot add person properties because given block is nil")
	}

	// Set first name if provided
	if firstName != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyFirstName, *firstName); err != nil {
			return fmt.Errorf("failed to set first name property: %w", err)
		}
	}

	// Set last name if provided
	if lastName != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyLastName, *lastName); err != nil {
			return fmt.Errorf("failed to set last name property: %w", err)
		}
	}

	// Set birthday if provided
	if birthday != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyBirthday, *birthday); err != nil {
			return fmt.Errorf("failed to set birthday property: %w", err)
		}
	}

	// Set relation type if provided
	if relationType != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyRelationType, *relationType); err != nil {
			return fmt.Errorf("failed to set relation type property: %w", err)
		}
	}

	// Set address if provided
	if address != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyAddress, *address); err != nil {
			return fmt.Errorf("failed to set address property: %w", err)
		}
	}

	// Set phone number if provided
	if phoneNumber != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyPhoneNumber, *phoneNumber); err != nil {
			return fmt.Errorf("failed to set phone number property: %w", err)
		}
	}

	// Set image URL if provided
	if imageURL != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyImageURL, *imageURL); err != nil {
			return fmt.Errorf("failed to set image URL property: %w", err)
		}
	}

	// Set description if provided
	if description != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyDescription, *description); err != nil {
			return fmt.Errorf("failed to set description property: %w", err)
		}
	}

	return nil
}

// RenderPersonProperties renders person properties in a human-readable format
func RenderPersonProperties(b Block) string {
	var result string

	// Get first name
	firstName, hasFirstName := b.Properties.GetString(PropertyKeyFirstName)

	// Get last name
	lastName, hasLastName := b.Properties.GetString(PropertyKeyLastName)

	// Get full name for the title
	var fullName string
	if hasFirstName && hasLastName {
		fullName = firstName + " " + lastName
	} else if hasFirstName {
		fullName = firstName
	} else if hasLastName {
		fullName = lastName
	} else {
		fullName = "Unnamed Person"
	}

	// Add the full name as the title
	result = fmt.Sprintf("# %s\n", fullName)

	// Get relation type
	relationType, hasRelationType := b.Properties.GetString(PropertyKeyRelationType)
	if hasRelationType && relationType != "" {
		result += fmt.Sprintf("**Relation:** %s\n", relationType)
	}

	// Get birthday
	birthday, hasBirthday := b.Properties.GetTime(PropertyKeyBirthday)
	if hasBirthday {
		age := calculateAge(birthday)
		result += fmt.Sprintf("**Birthday:** %s (Age: %d)\n", birthday.Format("January 2, 2006"), age)
	}

	// Get contact information
	var contactInfo []string

	// Get phone number
	phoneNumber, hasPhoneNumber := b.Properties.GetString(PropertyKeyPhoneNumber)
	if hasPhoneNumber && phoneNumber != "" {
		contactInfo = append(contactInfo, fmt.Sprintf("**Phone:** %s", phoneNumber))
	}

	// Get address
	address, hasAddress := b.Properties.GetString(PropertyKeyAddress)
	if hasAddress && address != "" {
		contactInfo = append(contactInfo, fmt.Sprintf("**Address:** %s", address))
	}

	// Add contact information
	if len(contactInfo) > 0 {
		result += "## Contact Information\n"
		result += strings.Join(contactInfo, "\n") + "\n"
	}

	// Get description
	description, hasDescription := b.Properties.GetString(PropertyKeyDescription)
	if hasDescription && description != "" {
		result += "## Notes\n"
		result += description + "\n"
	}

	return result
}

// calculateAge calculates the age based on a birthday
func calculateAge(birthday time.Time) int {
	now := time.Now()
	years := now.Year() - birthday.Year()

	// Adjust age if birthday hasn't occurred yet this year
	if now.YearDay() < birthday.YearDay() {
		years--
	}

	return years
}

// GetPersonProperties returns a list of all person property keys
func GetPersonProperties() []string {
	return []string{
		PropertyKeyFirstName,
		PropertyKeyLastName,
		PropertyKeyBirthday,
		PropertyKeyRelationType,
		PropertyKeyAddress,
		PropertyKeyPhoneNumber,
		PropertyKeyImageURL,
		PropertyKeyDescription,
	}
}
