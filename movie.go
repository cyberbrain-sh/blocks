package blocks

import (
	"fmt"
	"strconv"
	"strings"
)

// flattenArray recursively flattens nested arrays into a single flat array of strings
func flattenArray(arr []interface{}) []string {
	var result []string

	for _, item := range arr {
		switch v := item.(type) {
		case string:
			if v != "" {
				result = append(result, v)
			}
		case []interface{}:
			// Recursively flatten nested arrays
			nested := flattenArray(v)
			result = append(result, nested...)
		case []string:
			// Convert []string to []interface{} and flatten
			interfaceArr := make([]interface{}, len(v))
			for i, s := range v {
				interfaceArr[i] = s
			}
			nested := flattenArray(interfaceArr)
			result = append(result, nested...)
		default:
			// Convert any other type to string
			if str := fmt.Sprintf("%v", v); str != "" {
				result = append(result, str)
			}
		}
	}

	return result
}

func AddMoviePropertiesFromPage(block *Block, page PageData) error {
	if block == nil {
		return fmt.Errorf("cannot add movie properties because given block is nil")
	}

	// Extract basic page information
	title := page.GetTitle()
	description := page.GetDescription()
	imageURL := page.GetImage()
	originalURL := page.GetURL()

	// Extract custom metadata
	customMetadata := page.GetCustomMetadata()

	// Initialize pointers for AddMovieProperties
	var imdbID, tmdbID *string
	var releaseYear *int
	var rating, runtime, tagline, budget, revenue *string
	var genres, directors, cast *[]string
	var checked bool = false // Add checked property

	if customMetadata != nil {
		// Extract IDs
		if id, ok := customMetadata["imdb_id"]; ok && id != "" {
			imdbID = &id
		}

		if id, ok := customMetadata["tmdb_id"]; ok && id != "" {
			tmdbID = &id
		}

		// Extract release year if available
		if yearStr, ok := customMetadata["release_year"]; ok && yearStr != "Unknown" {
			if year, err := strconv.Atoi(yearStr); err == nil {
				releaseYear = &year
			}
		}

		// Extract other string fields
		if val, ok := customMetadata["rating"]; ok && val != "" {
			rating = &val
		}

		if val, ok := customMetadata["runtime"]; ok && val != "" {
			runtime = &val
		}

		if val, ok := customMetadata["tagline"]; ok && val != "" {
			tagline = &val
		}

		if val, ok := customMetadata["budget"]; ok && val != "" {
			budget = &val
		}

		if val, ok := customMetadata["revenue"]; ok && val != "" {
			revenue = &val
		}

		// Extract array fields by splitting on commas
		if genresStr, ok := customMetadata["genres"]; ok && genresStr != "" {
			genresList := strings.Split(genresStr, ", ")
			genres = &genresList
		}

		if directorsStr, ok := customMetadata["directors"]; ok && directorsStr != "" {
			directorsList := strings.Split(directorsStr, ", ")
			directors = &directorsList
		}

		if castStr, ok := customMetadata["cast"]; ok && castStr != "" {
			castList := strings.Split(castStr, ", ")
			cast = &castList
		}

		// Parse checked property if available
		if checkedStr, ok := customMetadata["checked"]; ok {
			checkedBool := checkedStr == "true"
			checked = checkedBool
		}
	}

	// Call AddMovieProperties with the extracted values
	return AddMovieProperties(block, &title, &description, &imageURL, &originalURL,
		imdbID, tmdbID, releaseYear, rating, runtime, tagline, budget, revenue,
		genres, directors, cast, &checked, true)
}

func AddMovieProperties(b *Block, title, description, imageURL, url, imdbID, tmdbID *string, releaseYear *int,
	rating, runtime, tagline, budget, revenue *string, genres, directors, cast *[]string, checked *bool, enriched bool) error {
	if b == nil {
		return fmt.Errorf("cannot add movie properties because given b is nil")
	}

	if checked == nil {
		pFalse := false
		checked = &pFalse
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

	// Movie specific properties
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

	if releaseYear != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyReleaseYear, *releaseYear); err != nil {
			return fmt.Errorf("failed to set release year property: %w", err)
		}
	}

	if rating != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyRating, *rating); err != nil {
			return fmt.Errorf("failed to set rating property: %w", err)
		}
	}

	if runtime != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyRuntime, *runtime); err != nil {
			return fmt.Errorf("failed to set runtime property: %w", err)
		}
	}

	if tagline != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTagline, *tagline); err != nil {
			return fmt.Errorf("failed to set tagline property: %w", err)
		}
	}

	if budget != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyBudget, *budget); err != nil {
			return fmt.Errorf("failed to set budget property: %w", err)
		}
	}

	if revenue != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyRevenue, *revenue); err != nil {
			return fmt.Errorf("failed to set revenue property: %w", err)
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

	if directors != nil && len(*directors) > 0 {
		directorInterfaces := make([]interface{}, len(*directors))
		for i, director := range *directors {
			directorInterfaces[i] = director
		}
		if err := b.Properties.ReplaceValue(PropertyKeyDirectors, directorInterfaces); err != nil {
			return fmt.Errorf("failed to set directors property: %w", err)
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

	if err := b.Properties.ReplaceValue(PropertyKeyEnriched, enriched); err != nil {
		return fmt.Errorf("failed to set enriched property: %w", err)
	}

	return nil
}

func RenderMovieProperties(b Block) string {
	// Get basic properties using proper type casting
	title, hasTitle := b.Properties.GetString(PropertyKeyTitle)
	url, hasURL := b.Properties.GetString(PropertyKeyURL)
	description, hasDesc := b.Properties.GetString(PropertyKeyDescription)

	// Get checked property
	checked, hasChecked := b.Properties.GetBool(PropertyKeyChecked)

	// Get movie-specific properties with proper type casting
	releaseYear, hasYear := b.Properties.GetInt(PropertyKeyReleaseYear)
	rating, hasRating := b.Properties.GetFloat(PropertyKeyRating)
	runtime, hasRuntime := b.Properties.GetInt(PropertyKeyRuntime)
	tagline, hasTagline := b.Properties.GetString(PropertyKeyTagline)

	// Get array properties and flatten them
	genresArr, hasGenres := b.Properties.GetArray(PropertyKeyGenres)
	directorsArr, hasDirectors := b.Properties.GetArray(PropertyKeyDirectors)
	castArr, hasCast := b.Properties.GetArray(PropertyKeyCast)

	// Flatten arrays
	var genres, directors, cast []string
	if hasGenres && len(genresArr) > 0 {
		genres = flattenArray(genresArr)
	}
	if hasDirectors && len(directorsArr) > 0 {
		directors = flattenArray(directorsArr)
	}
	if hasCast && len(castArr) > 0 {
		cast = flattenArray(castArr)
	}

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

	// Title with year and link
	if hasTitle && title != "" {
		titleText := title

		// Add year if available
		if hasYear && releaseYear > 0 {
			titleText = fmt.Sprintf("%s (%d)", titleText, releaseYear)
		}

		// Add URL if available
		if hasURL && url != "" {
			parts = append(parts, fmt.Sprintf("%s# [%s](%s)", prefix, titleText, url))
		} else {
			parts = append(parts, fmt.Sprintf("%s# %s", prefix, titleText))
		}
	}

	// Tagline
	if hasTagline && tagline != "" {
		parts = append(parts, fmt.Sprintf("*%s*", tagline))
	}

	// Rating and Runtime
	var stats []string
	if hasRating && rating > 0 {
		stats = append(stats, fmt.Sprintf("Rating: ⭐ %.1f", rating))
	}
	if hasRuntime && runtime > 0 {
		stats = append(stats, fmt.Sprintf("Runtime: %d min", runtime))
	}

	if len(stats) > 0 {
		parts = append(parts, fmt.Sprintf("**%s**", strings.Join(stats, " | ")))
	}

	// Genres
	if len(genres) > 0 {
		parts = append(parts, fmt.Sprintf("**Genres:** %s", strings.Join(genres, ", ")))
	}

	// Directors
	if len(directors) > 0 {
		parts = append(parts, fmt.Sprintf("**Directors:** %s", strings.Join(directors, ", ")))
	}

	// Cast
	if len(cast) > 0 {
		parts = append(parts, fmt.Sprintf("**Cast:** %s", strings.Join(cast, ", ")))
	}

	// Description
	if hasDesc && description != "" {
		parts = append(parts, fmt.Sprintf("\n%s", description))
	}

	return strings.Join(parts, "\n")
}

func GetMovieProperties() []string {
	return []string{
		PropertyKeyTitle,
		PropertyKeyDescription,
		PropertyKeyImageURL,
		PropertyKeyURL,
		PropertyKeyIMDBID,
		PropertyKeyTMDBID,
		PropertyKeyReleaseYear,
		PropertyKeyRating,
		PropertyKeyRuntime,
		PropertyKeyGenres,
		PropertyKeyDirectors,
		PropertyKeyCast,
		PropertyKeyTagline,
		PropertyKeyBudget,
		PropertyKeyRevenue,
		PropertyKeyEnriched,
		PropertyKeyChecked,
	}
}
