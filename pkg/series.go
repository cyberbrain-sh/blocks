package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

func AddSeriesPropertiesFromPage(b *Block, page PageData) error {
	if b == nil {
		return fmt.Errorf("cannot add series properties because given b is nil")
	}

	// Extract basic page information
	title := page.GetTitle()
	description := page.GetDescription()
	imageURL := page.GetImage()
	originalURL := page.GetURL()

	// Extract custom metadata
	customMetadata := page.GetCustomMetadata()

	// Initialize pointers for AddSeriesProperties
	var imdbID, tmdbID *string
	var firstAirYear, lastAirYear, numberOfSeasons, numberOfEpisodes *int
	var rating, status, seriesType *string
	var inProduction *bool
	var genres, creators, cast, networks *[]string
	var checked *bool = nil // Add checked property

	if customMetadata != nil {
		// Extract IDs
		if id, ok := customMetadata["imdb_id"]; ok && id != "" {
			imdbID = &id
		}

		if id, ok := customMetadata["tmdb_id"]; ok && id != "" {
			tmdbID = &id
		}

		// Extract first air year if available
		if yearStr, ok := customMetadata["first_air_year"]; ok && yearStr != "Unknown" {
			if year, err := strconv.Atoi(yearStr); err == nil {
				firstAirYear = &year
			}
		}

		// Extract last air year if available
		if yearStr, ok := customMetadata["last_air_year"]; ok && yearStr != "Unknown" {
			if year, err := strconv.Atoi(yearStr); err == nil {
				lastAirYear = &year
			}
		}

		// Extract other string fields
		if val, ok := customMetadata["rating"]; ok && val != "" {
			rating = &val
		}

		if val, ok := customMetadata["status"]; ok && val != "" {
			status = &val
		}

		if val, ok := customMetadata["type"]; ok && val != "" {
			seriesType = &val
		}

		// Parse in_production to boolean
		if inProductionStr, ok := customMetadata["in_production"]; ok {
			inProductionBool := inProductionStr == "true"
			inProduction = &inProductionBool
		}

		// Parse number of seasons to int
		if seasons, ok := customMetadata["number_of_seasons"]; ok && seasons != "" {
			if num, err := strconv.Atoi(seasons); err == nil {
				numberOfSeasons = &num
			}
		}

		// Parse number of episodes to int
		if episodes, ok := customMetadata["number_of_episodes"]; ok && episodes != "" {
			if num, err := strconv.Atoi(episodes); err == nil {
				numberOfEpisodes = &num
			}
		}

		// Extract array fields by splitting on commas
		if genresStr, ok := customMetadata["genres"]; ok && genresStr != "" {
			genresList := strings.Split(genresStr, ", ")
			genres = &genresList
		}

		if creatorsStr, ok := customMetadata["creators"]; ok && creatorsStr != "" {
			creatorsList := strings.Split(creatorsStr, ", ")
			creators = &creatorsList
		}

		if castStr, ok := customMetadata["cast"]; ok && castStr != "" {
			castList := strings.Split(castStr, ", ")
			cast = &castList
		}

		if networksStr, ok := customMetadata["networks"]; ok && networksStr != "" {
			networksList := strings.Split(networksStr, ", ")
			networks = &networksList
		}

		// Parse checked property if available
		if checkedStr, ok := customMetadata["checked"]; ok {
			checkedBool := checkedStr == "true"
			checked = &checkedBool
		}
	}

	// Call AddSeriesProperties with the extracted values
	return AddSeriesProperties(b, &title, &description, &imageURL, &originalURL,
		imdbID, tmdbID, firstAirYear, lastAirYear, numberOfSeasons, numberOfEpisodes,
		rating, status, seriesType, inProduction,
		genres, creators, cast, networks, checked, true)
}

func AddSeriesProperties(b *Block, title, description, imageURL, url, imdbID, tmdbID *string,
	firstAirYear, lastAirYear, numberOfSeasons, numberOfEpisodes *int,
	rating, status, seriesType *string, inProduction *bool,
	genres, creators, cast, networks *[]string, checked *bool, enriched bool) error {
	if b == nil {
		return fmt.Errorf("cannot add series properties because given b is nil")
	}

	// Common properties
	if title != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTitle, *title); err != nil {
			return fmt.Errorf("failed to set title property: %w", err)
		}
	}

	if description != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyDescription, *description); err != nil {
			return fmt.Errorf("failed to set description property: %w", err)
		}
	}

	if imageURL != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyImageURL, *imageURL); err != nil {
			return fmt.Errorf("failed to set image URL property: %w", err)
		}
	}

	if url != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyURL, *url); err != nil {
			return fmt.Errorf("failed to set URL property: %w", err)
		}
	}

	// Series specific properties
	if imdbID != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyIMDBID, *imdbID); err != nil {
			return fmt.Errorf("failed to set IMDB ID property: %w", err)
		}
	}

	if tmdbID != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTMDBID, *tmdbID); err != nil {
			return fmt.Errorf("failed to set TMDB ID property: %w", err)
		}
	}

	if firstAirYear != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyFirstAirYear, *firstAirYear); err != nil {
			return fmt.Errorf("failed to set first air year property: %w", err)
		}
	}

	if lastAirYear != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyLastAirYear, *lastAirYear); err != nil {
			return fmt.Errorf("failed to set last air year property: %w", err)
		}
	}

	if rating != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyRating, *rating); err != nil {
			return fmt.Errorf("failed to set rating property: %w", err)
		}
	}

	if status != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyStatus, *status); err != nil {
			return fmt.Errorf("failed to set status property: %w", err)
		}
	}

	if inProduction != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyInProduction, *inProduction); err != nil {
			return fmt.Errorf("failed to set in production property: %w", err)
		}
	}

	if numberOfSeasons != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyNumberOfSeasons, *numberOfSeasons); err != nil {
			return fmt.Errorf("failed to set number of seasons property: %w", err)
		}
	}

	if numberOfEpisodes != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyNumberOfEpisodes, *numberOfEpisodes); err != nil {
			return fmt.Errorf("failed to set number of episodes property: %w", err)
		}
	}

	if seriesType != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyType, *seriesType); err != nil {
			return fmt.Errorf("failed to set type property: %w", err)
		}
	}

	// Checked property
	if checked != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyChecked, *checked); err != nil {
			return fmt.Errorf("failed to set checked property: %w", err)
		}
	}

	// Array properties
	if genres != nil && len(*genres) > 0 {
		genreInterfaces := make([]interface{}, len(*genres))
		for i, genre := range *genres {
			genreInterfaces[i] = genre
		}
		if err := b.Properties.ReplaceValue(PropertyKeyGenres, genreInterfaces); err != nil {
			return fmt.Errorf("failed to set genres property: %w", err)
		}
	}

	if creators != nil && len(*creators) > 0 {
		creatorInterfaces := make([]interface{}, len(*creators))
		for i, creator := range *creators {
			creatorInterfaces[i] = creator
		}
		if err := b.Properties.ReplaceValue(PropertyKeyCreators, creatorInterfaces); err != nil {
			return fmt.Errorf("failed to set creators property: %w", err)
		}
	}

	if cast != nil && len(*cast) > 0 {
		castInterfaces := make([]interface{}, len(*cast))
		for i, actor := range *cast {
			castInterfaces[i] = actor
		}
		if err := b.Properties.ReplaceValue(PropertyKeyCast, castInterfaces); err != nil {
			return fmt.Errorf("failed to set cast property: %w", err)
		}
	}

	if networks != nil && len(*networks) > 0 {
		networkInterfaces := make([]interface{}, len(*networks))
		for i, network := range *networks {
			networkInterfaces[i] = network
		}
		if err := b.Properties.ReplaceValue(PropertyKeyNetworks, networkInterfaces); err != nil {
			return fmt.Errorf("failed to set networks property: %w", err)
		}
	}

	if err := b.Properties.ReplaceValue(PropertyKeyEnriched, enriched); err != nil {
		return fmt.Errorf("failed to set enriched property: %w", err)
	}

	return nil
}

func RenderSeriesProperties(b Block) string {
	// Get basic properties
	titleValue, hasTitle := b.Properties.Get(PropertyKeyTitle)
	title, titleOk := titleValue.(string)

	urlValue, hasURL := b.Properties.Get(PropertyKeyURL)
	url, urlOk := urlValue.(string)

	descValue, hasDesc := b.Properties.Get(PropertyKeyDescription)
	description, descOk := descValue.(string)

	// Get checked property
	checkedValue, hasChecked := b.Properties.Get(PropertyKeyChecked)
	checked, _ := checkedValue.(bool)

	// Get series-specific properties
	firstYearVal, hasFirstYear := b.Properties.Get(PropertyKeyFirstAirYear)
	lastYearVal, hasLastYear := b.Properties.Get(PropertyKeyLastAirYear)

	ratingVal, hasRating := b.Properties.Get(PropertyKeyRating)
	rating, ratingOk := ratingVal.(string)

	statusVal, hasStatus := b.Properties.Get(PropertyKeyStatus)
	status, statusOk := statusVal.(string)

	seasonsVal, hasSeasons := b.Properties.Get(PropertyKeyNumberOfSeasons)
	episodesVal, hasEpisodes := b.Properties.Get(PropertyKeyNumberOfEpisodes)

	// Array properties
	genresArr, hasGenres := b.Properties.GetArray(PropertyKeyGenres)
	creatorsArr, hasCreators := b.Properties.GetArray(PropertyKeyCreators)
	castArr, hasCast := b.Properties.GetArray(PropertyKeyCast)
	networksArr, hasNetworks := b.Properties.GetArray(PropertyKeyNetworks)

	// Build the markdown representation
	var parts []string

	// Add checkbox if checked property exists
	prefix := ""
	if hasChecked {
		if checked {
			prefix = "- [x] "
		} else {
			prefix = "- [ ] "
		}
	}

	// Title with year range and link
	if hasTitle && titleOk && title != "" {
		titleText := title

		// Add year range if available
		var firstYear, lastYear string
		if hasFirstYear {
			switch y := firstYearVal.(type) {
			case int:
				firstYear = strconv.Itoa(y)
			case float64:
				firstYear = strconv.Itoa(int(y))
			case string:
				firstYear = y
			}
		}

		if hasLastYear {
			switch y := lastYearVal.(type) {
			case int:
				lastYear = strconv.Itoa(y)
			case float64:
				lastYear = strconv.Itoa(int(y))
			case string:
				lastYear = y
			}
		}

		if firstYear != "" {
			if lastYear != "" && lastYear != firstYear {
				titleText = fmt.Sprintf("%s (%s-%s)", titleText, firstYear, lastYear)
			} else {
				titleText = fmt.Sprintf("%s (%s)", titleText, firstYear)
			}
		}

		// Add URL if available
		if hasURL && urlOk && url != "" {
			parts = append(parts, fmt.Sprintf("%s## [%s](%s)", prefix, titleText, url))
		} else {
			parts = append(parts, fmt.Sprintf("%s## %s", prefix, titleText))
		}
	}

	// Status
	if hasStatus && statusOk && status != "" {
		parts = append(parts, fmt.Sprintf("**Status:** %s", status))
	}

	// Rating and Seasons/Episodes
	var stats []string
	if hasRating && ratingOk && rating != "" {
		stats = append(stats, fmt.Sprintf("Rating: ⭐ %s", rating))
	}

	var seasonsEpisodes string
	if hasSeasons {
		var seasons string
		switch s := seasonsVal.(type) {
		case int:
			seasons = strconv.Itoa(s)
		case float64:
			seasons = strconv.Itoa(int(s))
		case string:
			seasons = s
		}

		if seasons != "" {
			seasonsEpisodes = fmt.Sprintf("%s Season", seasons)
			if seasons != "1" {
				seasonsEpisodes += "s"
			}
		}
	}

	if hasEpisodes {
		var episodes string
		switch e := episodesVal.(type) {
		case int:
			episodes = strconv.Itoa(e)
		case float64:
			episodes = strconv.Itoa(int(e))
		case string:
			episodes = e
		}

		if episodes != "" {
			if seasonsEpisodes != "" {
				seasonsEpisodes += fmt.Sprintf(", %s Episodes", episodes)
			} else {
				seasonsEpisodes = fmt.Sprintf("%s Episodes", episodes)
			}
		}
	}

	if seasonsEpisodes != "" {
		stats = append(stats, seasonsEpisodes)
	}

	if len(stats) > 0 {
		parts = append(parts, fmt.Sprintf("**%s**", strings.Join(stats, " | ")))
	}

	// Networks
	if hasNetworks && len(networksArr) > 0 {
		var networkStrs []string
		for _, n := range networksArr {
			if ns, ok := n.(string); ok && ns != "" {
				networkStrs = append(networkStrs, ns)
			}
		}

		if len(networkStrs) > 0 {
			parts = append(parts, fmt.Sprintf("**Networks:** %s", strings.Join(networkStrs, ", ")))
		}
	}

	// Genres
	if hasGenres && len(genresArr) > 0 {
		var genreStrs []string
		for _, g := range genresArr {
			if gs, ok := g.(string); ok && gs != "" {
				genreStrs = append(genreStrs, gs)
			}
		}

		if len(genreStrs) > 0 {
			parts = append(parts, fmt.Sprintf("**Genres:** %s", strings.Join(genreStrs, ", ")))
		}
	}

	// Creators
	if hasCreators && len(creatorsArr) > 0 {
		var creatorStrs []string
		for _, c := range creatorsArr {
			if cs, ok := c.(string); ok && cs != "" {
				creatorStrs = append(creatorStrs, cs)
			}
		}

		if len(creatorStrs) > 0 {
			parts = append(parts, fmt.Sprintf("**Creators:** %s", strings.Join(creatorStrs, ", ")))
		}
	}

	// Cast
	if hasCast && len(castArr) > 0 {
		var castStrs []string
		for _, c := range castArr {
			if cs, ok := c.(string); ok && cs != "" {
				castStrs = append(castStrs, cs)
			}
		}

		if len(castStrs) > 0 {
			parts = append(parts, fmt.Sprintf("**Cast:** %s", strings.Join(castStrs, ", ")))
		}
	}

	// Description
	if hasDesc && descOk && description != "" {
		parts = append(parts, fmt.Sprintf("\n%s", description))
	}

	return strings.Join(parts, "\n")
}

func GetSeriesProperties() []string {
	return []string{
		PropertyKeyTitle,
		PropertyKeyDescription,
		PropertyKeyImageURL,
		PropertyKeyURL,
		PropertyKeyIMDBID,
		PropertyKeyTMDBID,
		PropertyKeyFirstAirYear,
		PropertyKeyLastAirYear,
		PropertyKeyRating,
		PropertyKeyStatus,
		PropertyKeyInProduction,
		PropertyKeyNumberOfSeasons,
		PropertyKeyNumberOfEpisodes,
		PropertyKeyGenres,
		PropertyKeyCreators,
		PropertyKeyCast,
		PropertyKeyNetworks,
		PropertyKeyType,
		PropertyKeyEnriched,
		PropertyKeyChecked,
	}
}
