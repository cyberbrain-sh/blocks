package blocks

import (
	"fmt"
	"strings"
	"time"
)

// AddBookProperties adds book properties to the given block
func AddBookProperties(b *Block, title *string, description *string, imageURL *string, url *string,
	isbn *string, authors *[]string, publisher *string, publishedAt *time.Time,
	pageCount *int, genres *[]string, language *string, tagline *string, price *float64,
	authorBio *string, tableOfContents *string, reviews *[]string, enriched bool) error {

	if b == nil {
		return fmt.Errorf("cannot add book properties because given block is nil")
	}

	// Set title if provided
	if title != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTitle, *title); err != nil {
			return fmt.Errorf("failed to set title property: %w", err)
		}
	}

	// Set description if provided
	if description != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyDescription, *description); err != nil {
			return fmt.Errorf("failed to set description property: %w", err)
		}
	}

	// Set image URL if provided
	if imageURL != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyImageURL, *imageURL); err != nil {
			return fmt.Errorf("failed to set image URL property: %w", err)
		}
	}

	// Set URL if provided
	if url != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyURL, *url); err != nil {
			return fmt.Errorf("failed to set URL property: %w", err)
		}
	}

	// Set ISBN if provided
	if isbn != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyISBN, *isbn); err != nil {
			return fmt.Errorf("failed to set ISBN property: %w", err)
		}
	}

	// Set authors if provided
	if authors != nil {
		// Clear existing authors
		b.Properties.Delete(PropertyKeyAuthorName)

		// Add each author to the array
		for _, author := range *authors {
			if err := b.Properties.AppendValue(PropertyKeyAuthorName, author); err != nil {
				return fmt.Errorf("failed to add author: %w", err)
			}
		}
	}

	// Set publisher if provided
	if publisher != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyPublisher, *publisher); err != nil {
			return fmt.Errorf("failed to set publisher property: %w", err)
		}
	}

	// Set published at date if provided
	if publishedAt != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyPublishedAt, *publishedAt); err != nil {
			return fmt.Errorf("failed to set published date property: %w", err)
		}
	}

	// Set page count if provided
	if pageCount != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyPageCount, *pageCount); err != nil {
			return fmt.Errorf("failed to set page count property: %w", err)
		}
	}

	// Set genres if provided
	if genres != nil {
		// Clear existing genres
		b.Properties.Delete(PropertyKeyGenres)

		// Add each genre to the array
		for _, genre := range *genres {
			if err := b.Properties.AppendValue(PropertyKeyGenres, genre); err != nil {
				return fmt.Errorf("failed to add genre: %w", err)
			}
		}
	}

	// Set language if provided
	if language != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyLanguage, *language); err != nil {
			return fmt.Errorf("failed to set language property: %w", err)
		}
	}

	// Set tagline if provided
	if tagline != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTagline, *tagline); err != nil {
			return fmt.Errorf("failed to set tagline property: %w", err)
		}
	}

	// Set price if provided
	if price != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyPrice, *price); err != nil {
			return fmt.Errorf("failed to set price property: %w", err)
		}
	}

	// Set author bio if provided
	if authorBio != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyAuthorBio, *authorBio); err != nil {
			return fmt.Errorf("failed to set author bio property: %w", err)
		}
	}

	// Set table of contents if provided
	if tableOfContents != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTableOfContents, *tableOfContents); err != nil {
			return fmt.Errorf("failed to set table of contents property: %w", err)
		}
	}

	// Set reviews if provided
	if reviews != nil {
		// Clear existing reviews
		b.Properties.Delete(PropertyKeyReviews)

		// Add each review to the array
		for _, review := range *reviews {
			if err := b.Properties.AppendValue(PropertyKeyReviews, review); err != nil {
				return fmt.Errorf("failed to add review: %w", err)
			}
		}
	}

	// Set enriched status
	if err := b.Properties.ReplaceValue(PropertyKeyEnriched, enriched); err != nil {
		return fmt.Errorf("failed to set enriched property: %w", err)
	}

	return nil
}

// RenderBookProperties renders book properties in a human-readable format
func RenderBookProperties(b Block) string {
	var result string

	// Get title
	title, hasTitle := b.Properties.GetString(PropertyKeyTitle)

	// Get authors
	authors, hasAuthors := b.Properties.GetStringArray(PropertyKeyAuthorName)

	// Get publication date
	publicationDate, hasPublicationDate := b.Properties.GetTime(PropertyKeyPublishedAt)

	// Get ISBN
	isbn, hasISBN := b.Properties.GetString(PropertyKeyISBN)

	// Get page count
	pageCount, hasPageCount := b.Properties.GetInt(PropertyKeyPageCount)

	// Get publisher
	publisher, hasPublisher := b.Properties.GetString(PropertyKeyPublisher)

	// Get genres
	genres, hasGenres := b.Properties.GetStringArray(PropertyKeyGenres)

	// Get language
	language, hasLanguage := b.Properties.GetString(PropertyKeyLanguage)

	// Get tagline
	tagline, hasTagline := b.Properties.GetString(PropertyKeyTagline)

	// Get price
	price, hasPrice := b.Properties.GetFloat(PropertyKeyPrice)

	// Get author bio
	authorBio, hasAuthorBio := b.Properties.GetString(PropertyKeyAuthorBio)

	// Get table of contents
	tableOfContents, hasTableOfContents := b.Properties.GetString(PropertyKeyTableOfContents)

	// Get reviews
	reviews, hasReviews := b.Properties.GetStringArray(PropertyKeyReviews)

	// Format the book title with publication year if available
	if hasTitle {
		result = fmt.Sprintf("# %s", title)

		if hasPublicationDate {
			result += fmt.Sprintf(" (%d)", publicationDate.Year())
		}
		result += "\n"
	}

	// Add tagline if available
	if hasTagline && tagline != "" {
		result += fmt.Sprintf("*%s*\n", tagline)
	}

	// Add authors if available
	if hasAuthors && len(authors) > 0 {
		result += "**Authors:** "
		for i, author := range authors {
			if i > 0 {
				result += ", "
			}
			result += author
		}
		result += "\n"
	}

	// Add additional details in a structured format
	var details []string

	if hasISBN {
		details = append(details, fmt.Sprintf("**ISBN:** %s", isbn))
	}

	if hasPublisher {
		details = append(details, fmt.Sprintf("**Publisher:** %s", publisher))
	}

	if hasLanguage {
		details = append(details, fmt.Sprintf("**Language:** %s", language))
	}

	if hasPageCount {
		details = append(details, fmt.Sprintf("**Pages:** %d", pageCount))
	}

	if hasPrice {
		details = append(details, fmt.Sprintf("**Price:** $%.2f", price))
	}

	// Join details with separator
	if len(details) > 0 {
		result += strings.Join(details, " | ") + "\n"
	}

	// Add genres if available
	if hasGenres && len(genres) > 0 {
		result += "**Genres:** "
		for i, genre := range genres {
			if i > 0 {
				result += ", "
			}
			result += genre
		}
		result += "\n"
	}

	// Add description if available
	description, hasDescription := b.Properties.GetString(PropertyKeyDescription)
	if hasDescription && description != "" {
		result += "---\n" + description + "\n"
	}

	// Add author bio if available
	if hasAuthorBio && authorBio != "" {
		result += "\n**Author Bio:**\n" + authorBio + "\n"
	}

	// Add table of contents if available
	if hasTableOfContents && tableOfContents != "" {
		result += "\n**Table of Contents:**\n" + tableOfContents + "\n"
	}

	// Add reviews if available
	if hasReviews && len(reviews) > 0 {
		result += "\n**Reviews:**\n"
		for _, review := range reviews {
			result += "- " + review + "\n"
		}
	}

	return result
}

// GetBookProperties returns a list of all book property keys
func GetBookProperties() []string {
	return []string{
		PropertyKeyTitle,
		PropertyKeyDescription,
		PropertyKeyImageURL,
		PropertyKeyURL,
		PropertyKeyISBN,
		PropertyKeyAuthorName,
		PropertyKeyPublisher,
		PropertyKeyPublishedAt,
		PropertyKeyPageCount,
		PropertyKeyGenres,
		PropertyKeyLanguage,
		PropertyKeyTagline,
		PropertyKeyPrice,
		PropertyKeyAuthorBio,
		PropertyKeyTableOfContents,
		PropertyKeyReviews,
		PropertyKeyEnriched,
	}
}
