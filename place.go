package blocks

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// PlaceData represents the structure of place data in JSON format
type PlaceData struct {
	Name        string    `json:"name"`
	PlaceType   string    `json:"place_type,omitempty"`
	Coordinates []float64 `json:"coordinates,omitempty"` // [latitude, longitude]
	MapURL      string    `json:"map_url,omitempty"`
	Address     string    `json:"address,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Rating      float64   `json:"rating,omitempty"`
	VisitedDate string    `json:"visited_date,omitempty"` // ISO 8601 format
	URL         string    `json:"url,omitempty"`
	ImageURL    string    `json:"image_url,omitempty"`
	Description string    `json:"description,omitempty"`
	Reviews     []string  `json:"reviews,omitempty"`
}

// AddPlacePropertiesFromJSON parses a json.RawMessage into a PlaceData struct
// and adds the place properties to the given block
func AddPlacePropertiesFromJSON(block *Block, rawJSON json.RawMessage) error {
	if block == nil {
		return fmt.Errorf("cannot add place properties because given block is nil")
	}

	// Parse the JSON data into a PlaceData struct
	var placeData PlaceData
	if err := json.Unmarshal(rawJSON, &placeData); err != nil {
		return fmt.Errorf("failed to unmarshal place data: %w", err)
	}

	// Validate required fields
	if placeData.Name == "" {
		return fmt.Errorf("place data must include a name")
	}

	// Convert struct fields to pointers for AddPlaceProperties
	name := placeData.Name

	var placeType, mapURL, address, phoneNumber, url, imageURL, description *string
	var rating *float64
	var coordinates *[]float64
	var visitedDate *time.Time
	var reviews *[]string

	if placeData.PlaceType != "" {
		placeType = &placeData.PlaceType
	}

	if len(placeData.Coordinates) == 2 {
		coordinates = &placeData.Coordinates
	}

	if placeData.MapURL != "" {
		mapURL = &placeData.MapURL
	}

	if placeData.Address != "" {
		address = &placeData.Address
	}

	if placeData.PhoneNumber != "" {
		phoneNumber = &placeData.PhoneNumber
	}

	if placeData.Rating > 0 {
		rating = &placeData.Rating
	}

	if placeData.VisitedDate != "" {
		// Parse the date string into a time.Time object
		parsedTime, err := time.Parse(time.RFC3339, placeData.VisitedDate)
		if err == nil {
			visitedDate = &parsedTime
		}
	}

	if placeData.URL != "" {
		url = &placeData.URL
	}

	if placeData.ImageURL != "" {
		imageURL = &placeData.ImageURL
	}

	if placeData.Description != "" {
		description = &placeData.Description
	}

	if len(placeData.Reviews) > 0 {
		reviews = &placeData.Reviews
	}

	// Call AddPlaceProperties with the extracted data
	return AddPlaceProperties(block, &name, placeType, coordinates, mapURL, address,
		phoneNumber, rating, visitedDate, url, imageURL, description, reviews)
}

// AddPlaceProperties adds place properties to the given block
func AddPlaceProperties(b *Block, name *string, placeType *string, coordinates *[]float64,
	mapURL *string, address *string, phoneNumber *string, rating *float64,
	visitedDate *time.Time, url *string, imageURL *string, description *string,
	reviews *[]string) error {

	if b == nil {
		return fmt.Errorf("cannot add place properties because given block is nil")
	}

	// Set place name if provided
	if name != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTitle, *name); err != nil {
			return fmt.Errorf("failed to set title property: %w", err)
		}
	}

	// Set place type if provided
	if placeType != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyPlaceType, *placeType); err != nil {
			return fmt.Errorf("failed to set place type property: %w", err)
		}
	}

	// Set coordinates if provided - must be an array of [latitude, longitude]
	if coordinates != nil {
		// Validate coordinates
		if len(*coordinates) != 2 {
			return fmt.Errorf("coordinates must be an array of exactly 2 values [latitude, longitude], got %d values", len(*coordinates))
		}

		if err := b.Properties.ReplaceValue(PropertyKeyCoordinates, *coordinates); err != nil {
			return fmt.Errorf("failed to set coordinates property: %w", err)
		}
	}

	// Set map URL if provided
	if mapURL != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyMapURL, *mapURL); err != nil {
			return fmt.Errorf("failed to set map URL property: %w", err)
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

	// Set rating if provided
	if rating != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyRating, *rating); err != nil {
			return fmt.Errorf("failed to set rating property: %w", err)
		}
	}

	// Set visited date if provided
	if visitedDate != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyVisitedDate, *visitedDate); err != nil {
			return fmt.Errorf("failed to set visited date property: %w", err)
		}
	}

	// Set URL if provided
	if url != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyURL, *url); err != nil {
			return fmt.Errorf("failed to set URL property: %w", err)
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

	// Set reviews if provided
	if reviews != nil {
		// Clear existing reviews
		b.Properties.Delete(PropertyKeyPlaceReviews)

		// Add each review to the array
		for _, review := range *reviews {
			if err := b.Properties.AppendValue(PropertyKeyPlaceReviews, review); err != nil {
				return fmt.Errorf("failed to add review: %w", err)
			}
		}
	}

	return nil
}

// RenderPlaceProperties renders place properties in a human-readable format
func RenderPlaceProperties(b Block) string {
	var result string

	// Get place name
	name, hasName := b.Properties.GetString(PropertyKeyTitle)

	// Use name as the title if available
	if hasName && name != "" {
		result = fmt.Sprintf("# %s\n", name)
	} else {
		result = "# Unnamed Place\n"
	}

	// Get place type and add as subtitle
	placeType, hasPlaceType := b.Properties.GetString(PropertyKeyPlaceType)
	if hasPlaceType && placeType != "" {
		result += fmt.Sprintf("*%s*\n", placeType)
	}

	// Get URL
	url, hasURL := b.Properties.GetString(PropertyKeyURL)
	if hasURL && url != "" {
		result += fmt.Sprintf("[Website](%s)\n", url)
	}

	// Get rating
	rating, hasRating := b.Properties.GetFloat(PropertyKeyRating)
	if hasRating {
		// Format rating with stars
		stars := ""
		for i := 0; i < int(rating); i++ {
			stars += "⭐"
		}

		result += fmt.Sprintf("**Rating:** %s (%.1f/5)\n", stars, rating)
	}

	// Get coordinates
	var coordsString string
	coordsArr, hasCoords := b.Properties.GetArray(PropertyKeyCoordinates)
	if hasCoords && len(coordsArr) == 2 {
		lat, latOk := coordsArr[0].(float64)
		lng, lngOk := coordsArr[1].(float64)

		if latOk && lngOk {
			coordsString = fmt.Sprintf("%.6f, %.6f", lat, lng)
		}
	}

	// Get map URL
	mapURL, hasMapURL := b.Properties.GetString(PropertyKeyMapURL)

	// Add location information
	var locationInfo []string

	// Get address
	address, hasAddress := b.Properties.GetString(PropertyKeyAddress)
	if hasAddress && address != "" {
		locationInfo = append(locationInfo, fmt.Sprintf("**Address:** %s", address))
	}

	// Add coordinates if available
	if coordsString != "" {
		// If map URL is available, make the coordinates a link
		if hasMapURL && mapURL != "" {
			locationInfo = append(locationInfo, fmt.Sprintf("**Coordinates:** [%s](%s)", coordsString, mapURL))
		} else {
			locationInfo = append(locationInfo, fmt.Sprintf("**Coordinates:** %s", coordsString))
		}
	} else if hasMapURL && mapURL != "" {
		// If only map URL is available, add a direct link
		locationInfo = append(locationInfo, fmt.Sprintf("**Map:** [Open in Google Maps](%s)", mapURL))
	}

	// Get phone number
	phoneNumber, hasPhoneNumber := b.Properties.GetString(PropertyKeyPhoneNumber)
	if hasPhoneNumber && phoneNumber != "" {
		locationInfo = append(locationInfo, fmt.Sprintf("**Phone:** %s", phoneNumber))
	}

	// Add location information section
	if len(locationInfo) > 0 {
		result += "\n## Location & Contact Information\n"
		result += strings.Join(locationInfo, "\n") + "\n"
	}

	// Get visited date
	visitedDate, hasVisitedDate := b.Properties.GetTime(PropertyKeyVisitedDate)
	if hasVisitedDate {
		result += fmt.Sprintf("\n**Last Visited:** %s\n", visitedDate.Format("January 2, 2006"))
	}

	// Get description
	description, hasDescription := b.Properties.GetString(PropertyKeyDescription)
	if hasDescription && description != "" {
		result += "\n## Description\n"
		result += description + "\n"
	}

	// Get reviews
	reviews, hasReviews := b.Properties.GetStringArray(PropertyKeyPlaceReviews)
	if hasReviews && len(reviews) > 0 {
		result += "\n## Reviews\n"
		for _, review := range reviews {
			result += fmt.Sprintf("- \"%s\"\n", review)
		}
	}

	return result
}

// GetPlaceProperties returns a list of all place property keys
func GetPlaceProperties() []string {
	return []string{
		PropertyKeyTitle,
		PropertyKeyPlaceType,
		PropertyKeyCoordinates,
		PropertyKeyMapURL,
		PropertyKeyAddress,
		PropertyKeyPhoneNumber,
		PropertyKeyRating,
		PropertyKeyVisitedDate,
		PropertyKeyURL,
		PropertyKeyImageURL,
		PropertyKeyDescription,
		PropertyKeyPlaceReviews,
	}
}

// GenerateGoogleMapsURL creates a Google Maps URL from coordinates
func GenerateGoogleMapsURL(latitude, longitude float64) string {
	return fmt.Sprintf("https://www.google.com/maps?q=%f,%f", latitude, longitude)
}
