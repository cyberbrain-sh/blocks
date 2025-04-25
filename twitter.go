package blocks

import (
	"fmt"
	"strconv"
)

func AddTwitterPropertiesFromPage(block *Block, page PageData) error {
	if block == nil {
		return fmt.Errorf("cannot add Twitter properties because given block is nil")
	}

	// Extract basic page information
	title := page.GetTitle()
	description := page.GetDescription()
	imageURL := page.GetImage()
	originalURL := page.GetURL()

	// Extract custom metadata
	customMetadata := page.GetCustomMetadata()

	// Initialize Twitter-specific properties
	var tweetID, username, authorName *string
	var likesCount, retweetsCount, repliesCount *int
	var tweetedAt *string

	if customMetadata != nil {
		// Extract Twitter-specific properties from custom metadata
		if id, ok := customMetadata["tweet_id"]; ok && id != "" {
			tweetID = &id
		}

		if user, ok := customMetadata["username"]; ok && user != "" {
			username = &user
		}

		if author, ok := customMetadata["author_name"]; ok && author != "" {
			authorName = &author
		}

		if date, ok := customMetadata["tweeted_at"]; ok && date != "" {
			tweetedAt = &date
		}

		// Process numeric properties
		if likesStr, ok := customMetadata["likes_count"]; ok && likesStr != "" {
			if likes, err := strconv.Atoi(likesStr); err == nil {
				likesCount = &likes
			}
		}

		if retweetsStr, ok := customMetadata["retweets_count"]; ok && retweetsStr != "" {
			if retweets, err := strconv.Atoi(retweetsStr); err == nil {
				retweetsCount = &retweets
			}
		}

		if repliesStr, ok := customMetadata["replies_count"]; ok && repliesStr != "" {
			if replies, err := strconv.Atoi(repliesStr); err == nil {
				repliesCount = &replies
			}
		}
	}

	// Call AddTwitterProperties with the extracted values
	return AddTwitterProperties(block, &originalURL, &title, &description, &imageURL,
		tweetID, username, authorName, tweetedAt, likesCount, retweetsCount, repliesCount, true)
}

func AddTwitterProperties(b *Block, url, title, description, urlImage,
	tweetID, username, authorName, tweetedAt *string,
	likesCount, retweetsCount, repliesCount *int, enriched bool) error {
	if b == nil {
		return fmt.Errorf("cannot add Twitter properties because given b is nil")
	}

	// Set common properties
	if url != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyURL, *url); err != nil {
			return fmt.Errorf("failed to set URL property: %w", err)
		}
	}

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

	if urlImage != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyImageURL, *urlImage); err != nil {
			return fmt.Errorf("failed to set image URL property: %w", err)
		}
	}

	// Set Twitter-specific properties
	if tweetID != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyTweetID, *tweetID); err != nil {
			return fmt.Errorf("failed to set tweet ID property: %w", err)
		}
	}

	if username != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyUsername, *username); err != nil {
			return fmt.Errorf("failed to set username property: %w", err)
		}
	}

	if authorName != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyAuthorName, *authorName); err != nil {
			return fmt.Errorf("failed to set author name property: %w", err)
		}
	}

	if tweetedAt != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyPublishedAt, *tweetedAt); err != nil {
			return fmt.Errorf("failed to set tweeted at property: %w", err)
		}
	}

	if likesCount != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyLikeCount, *likesCount); err != nil {
			return fmt.Errorf("failed to set likes count property: %w", err)
		}
	}

	if retweetsCount != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyRetweetCount, *retweetsCount); err != nil {
			return fmt.Errorf("failed to set retweets count property: %w", err)
		}
	}

	if repliesCount != nil {
		if err := b.Properties.ReplaceValue(PropertyKeyCommentCount, *repliesCount); err != nil {
			return fmt.Errorf("failed to set replies count property: %w", err)
		}
	}

	if err := b.Properties.ReplaceValue(PropertyKeyEnriched, enriched); err != nil {
		return fmt.Errorf("failed to set enriched property: %w", err)
	}

	return nil
}

func RenderTwitterProperties(b Block) string {
	// Get basic properties
	urlValue, hasURL := b.Properties.Get(PropertyKeyURL)
	url, urlOk := urlValue.(string)

	titleValue, hasTitle := b.Properties.Get(PropertyKeyTitle)
	title, titleOk := titleValue.(string)

	// Get Twitter-specific properties
	usernameValue, hasUsername := b.Properties.Get(PropertyKeyUsername)
	username, usernameOk := usernameValue.(string)

	authorNameValue, hasAuthorName := b.Properties.Get(PropertyKeyAuthorName)
	authorName, authorNameOk := authorNameValue.(string)

	tweetedAtValue, hasTweetedAt := b.Properties.Get(PropertyKeyPublishedAt)
	tweetedAt, tweetedAtOk := tweetedAtValue.(string)

	// Get engagement metrics
	likesValue, hasLikes := b.Properties.Get(PropertyKeyLikeCount)
	retweetsValue, hasRetweets := b.Properties.Get(PropertyKeyRetweetCount)
	repliesValue, hasReplies := b.Properties.Get(PropertyKeyCommentCount)

	// Format as markdown with Twitter information
	var result string

	// Start with author information
	if hasUsername && usernameOk && username != "" {
		if hasAuthorName && authorNameOk && authorName != "" {
			result = fmt.Sprintf("**%s** (@%s)", authorName, username)
		} else {
			result = fmt.Sprintf("@%s", username)
		}

		// Add tweet timestamp if available
		if hasTweetedAt && tweetedAtOk && tweetedAt != "" {
			result += fmt.Sprintf(" • %s", tweetedAt)
		}

		result += "\n\n"
	}

	// Add the tweet content
	if hasTitle && titleOk && title != "" {
		result += title + "\n\n"
	}

	// Add engagement metrics if available
	var metrics []string
	if hasLikes {
		likes, ok := likesValue.(int)
		if ok {
			metrics = append(metrics, fmt.Sprintf("❤️ %d", likes))
		}
	}

	if hasRetweets {
		retweets, ok := retweetsValue.(int)
		if ok {
			metrics = append(metrics, fmt.Sprintf("🔄 %d", retweets))
		}
	}

	if hasReplies {
		replies, ok := repliesValue.(int)
		if ok {
			metrics = append(metrics, fmt.Sprintf("💬 %d", replies))
		}
	}

	if len(metrics) > 0 {
		result += fmt.Sprintf("%s\n\n", metrics)
	}

	// Add URL as a link at the end
	if hasURL && urlOk && url != "" {
		result += fmt.Sprintf("[View on Twitter](%s)", url)
	}

	return result
}

func GetTwitterProperties() []string {
	return []string{
		PropertyKeyTitle,        // Tweet text
		PropertyKeyURL,          // Tweet URL
		PropertyKeyDescription,  // Additional context/description
		PropertyKeyImageURL,     // Tweet image if any
		PropertyKeyTweetID,      // Unique tweet ID
		PropertyKeyUsername,     // Twitter handle (@username)
		PropertyKeyAuthorName,   // Display name of the author
		PropertyKeyPublishedAt,  // When the tweet was posted
		PropertyKeyLikeCount,    // Number of likes
		PropertyKeyRetweetCount, // Number of retweets
		PropertyKeyCommentCount, // Number of replies
		PropertyKeyEnriched,
	}
}
