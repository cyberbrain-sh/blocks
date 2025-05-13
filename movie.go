package blocks

import (
	"fmt"
	"strconv"
	"strings"
)

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

	// Get movie-specific properties
	yearVal, hasYear := b.Properties.Get(PropertyKeyReleaseYear)

	ratingVal, hasRating := b.Properties.Get(PropertyKeyRating)
	rating, ratingOk := ratingVal.(string)

	runtimeVal, hasRuntime := b.Properties.Get(PropertyKeyRuntime)
	runtime, runtimeOk := runtimeVal.(string)

	taglineVal, hasTagline := b.Properties.Get(PropertyKeyTagline)
	tagline, taglineOk := taglineVal.(string)

	// Array properties
	genresArr, hasGenres := b.Properties.GetArray(PropertyKeyGenres)
	directorsArr, hasDirectors := b.Properties.GetArray(PropertyKeyDirectors)
	castArr, hasCast := b.Properties.GetArray(PropertyKeyCast)

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
	if hasTitle && titleOk && title != "" {
		titleText := title

		// Add year if available
		if hasYear {
			var year string
			switch y := yearVal.(type) {
			case int:
				year = strconv.Itoa(y)
			case float64:
				year = strconv.Itoa(int(y))
			case string:
				year = y
			}

			if year != "" {
				titleText = fmt.Sprintf("%s (%s)", titleText, year)
			}
		}

		// Add URL if available
		if hasURL && urlOk && url != "" {
			parts = append(parts, fmt.Sprintf("%s# [%s](%s)", prefix, titleText, url))
		} else {
			parts = append(parts, fmt.Sprintf("%s# %s", prefix, titleText))
		}
	}

	// Tagline
	if hasTagline && taglineOk && tagline != "" {
		parts = append(parts, fmt.Sprintf("*%s*", tagline))
	}

	// Rating and Runtime
	var stats []string
	if hasRating && ratingOk && rating != "" {
		stats = append(stats, fmt.Sprintf("Rating: ⭐ %s", rating))
	}
	if hasRuntime && runtimeOk && runtime != "" {
		stats = append(stats, fmt.Sprintf("Runtime: %s", runtime))
	}

	if len(stats) > 0 {
		parts = append(parts, fmt.Sprintf("**%s**", strings.Join(stats, " | ")))
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

	// Directors
	if hasDirectors && len(directorsArr) > 0 {
		var directorStrs []string
		for _, d := range directorsArr {
			if ds, ok := d.(string); ok && ds != "" {
				directorStrs = append(directorStrs, ds)
			}
		}

		if len(directorStrs) > 0 {
			parts = append(parts, fmt.Sprintf("**Directors:** %s", strings.Join(directorStrs, ", ")))
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
