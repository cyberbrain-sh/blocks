package blocks

import (
	"encoding/json"
	"fmt"
	"time"

	structured "github.com/cyberbrain-sh/blocks/data_helper"
)

// AddTweetPropertiesFromDataTweet adds properties to a block from a structured DataTweet
func AddTweetPropertiesFromDataTweet(block *Block, tweet *structured.DataTweet) error {
	if block == nil {
		return fmt.Errorf("cannot add Tweet properties because given block is nil")
	}

	if tweet == nil {
		return fmt.Errorf("cannot add Tweet properties because given tweet is nil")
	}

	// Convert CreatedAt time to string
	tweetedAt := tweet.CreatedAt.Format(time.RFC3339)

	// Set basic properties
	if err := block.Properties.ReplaceValue(PropertyKeyURL, tweet.URL); err != nil {
		return fmt.Errorf("failed to set URL property: %w", err)
	}

	if err := block.Properties.ReplaceValue(PropertyKeyTitle, tweet.Text); err != nil {
		return fmt.Errorf("failed to set title property: %w", err)
	}

	if err := block.Properties.ReplaceValue(PropertyKeyTweetID, tweet.ID); err != nil {
		return fmt.Errorf("failed to set tweet ID property: %w", err)
	}

	if err := block.Properties.ReplaceValue(PropertyKeyAuthorID, tweet.AuthorID); err != nil {
		return fmt.Errorf("failed to set author ID property: %w", err)
	}

	// Set author information
	if tweet.Username != "" {
		if err := block.Properties.ReplaceValue(PropertyKeyUsername, tweet.Username); err != nil {
			return fmt.Errorf("failed to set username property: %w", err)
		}
	}

	if tweet.AuthorName != "" {
		if err := block.Properties.ReplaceValue(PropertyKeyAuthorName, tweet.AuthorName); err != nil {
			return fmt.Errorf("failed to set author name property: %w", err)
		}
	}

	// Set timing information
	if err := block.Properties.ReplaceValue(PropertyKeyPublishedAt, tweetedAt); err != nil {
		return fmt.Errorf("failed to set published at property: %w", err)
	}

	// Set engagement metrics
	if err := block.Properties.ReplaceValue(PropertyKeyLikeCount, tweet.LikeCount); err != nil {
		return fmt.Errorf("failed to set like count property: %w", err)
	}

	if err := block.Properties.ReplaceValue(PropertyKeyRetweetCount, tweet.RetweetCount); err != nil {
		return fmt.Errorf("failed to set retweet count property: %w", err)
	}

	if err := block.Properties.ReplaceValue(PropertyKeyCommentCount, tweet.CommentCount); err != nil {
		return fmt.Errorf("failed to set comment count property: %w", err)
	}

	if tweet.QuoteCount > 0 {
		if err := block.Properties.ReplaceValue(PropertyKeyQuoteCount, tweet.QuoteCount); err != nil {
			return fmt.Errorf("failed to set quote count property: %w", err)
		}
	}

	// Set image URL from main image if available
	if tweet.ImageURL != "" {
		if err := block.Properties.ReplaceValue(PropertyKeyImageURL, tweet.ImageURL); err != nil {
			return fmt.Errorf("failed to set image URL property: %w", err)
		}
	}

	// Set tweet metadata
	if tweet.ConversationID != "" {
		if err := block.Properties.ReplaceValue(PropertyKeyConversationID, tweet.ConversationID); err != nil {
			return fmt.Errorf("failed to set conversation ID property: %w", err)
		}
	}

	if tweet.Language != "" {
		if err := block.Properties.ReplaceValue(PropertyKeyLanguage, tweet.Language); err != nil {
			return fmt.Errorf("failed to set language property: %w", err)
		}
	}

	if tweet.Source != "" {
		if err := block.Properties.ReplaceValue(PropertyKeySource, tweet.Source); err != nil {
			return fmt.Errorf("failed to set source property: %w", err)
		}
	}

	// Set media-related properties
	if err := block.Properties.ReplaceValue(PropertyKeyHasMedia, tweet.HasMedia); err != nil {
		return fmt.Errorf("failed to set has media property: %w", err)
	}

	if tweet.MediaCount > 0 {
		if err := block.Properties.ReplaceValue(PropertyKeyMediaCount, tweet.MediaCount); err != nil {
			return fmt.Errorf("failed to set media count property: %w", err)
		}
	}

	// Reset media URLs array field with empty array
	if err := block.Properties.ReplaceValue(PropertyKeyMediaURLs, []string{}); err != nil {
		return fmt.Errorf("failed to reset media URLs property: %w", err)
	}

	// Append media URLs one by one if available
	if len(tweet.MediaURLs) > 0 {
		for _, mediaURL := range tweet.MediaURLs {
			if err := block.Properties.AppendValue(PropertyKeyMediaURLs, mediaURL); err != nil {
				return fmt.Errorf("failed to append media URL: %w", err)
			}
		}
	}

	// Reset media info array field with empty array
	block.Properties.Delete(PropertyKeyMediaInfo)

	// Store media details as strings if available
	if len(tweet.Media) > 0 {
		for _, media := range tweet.Media {
			mediaInfo := fmt.Sprintf("%s:%s", media.MediaKey, media.URL)
			if err := block.Properties.AppendValue(PropertyKeyMediaInfo, mediaInfo); err != nil {
				return fmt.Errorf("failed to append media info: %w", err)
			}
		}
	}

	// Reset external URLs array field with empty array
	block.Properties.Delete(PropertyKeyExternalURLs)

	// Store external URLs if available
	if len(tweet.ExternalURLs) > 0 {
		for _, urlInfo := range tweet.ExternalURLs {
			if err := block.Properties.AppendValue(PropertyKeyExternalURLs, urlInfo.ExpandedURL); err != nil {
				return fmt.Errorf("failed to append external URL: %w", err)
			}
		}
	}

	// Set tweet type flags
	if err := block.Properties.ReplaceValue(PropertyKeyIsRetweet, tweet.IsRetweet); err != nil {
		return fmt.Errorf("failed to set is retweet property: %w", err)
	}

	if err := block.Properties.ReplaceValue(PropertyKeyIsReply, tweet.IsReply); err != nil {
		return fmt.Errorf("failed to set is reply property: %w", err)
	}

	if err := block.Properties.ReplaceValue(PropertyKeyIsQuote, tweet.IsQuote); err != nil {
		return fmt.Errorf("failed to set is quote property: %w", err)
	}

	// Set the enriched flag
	if err := block.Properties.ReplaceValue(PropertyKeyEnriched, true); err != nil {
		return fmt.Errorf("failed to set enriched property: %w", err)
	}

	return nil
}

// AddTweetProperties adds tweet properties to a block
func AddTweetProperties(b *Block, url, title, description, urlImage,
	tweetID, username, authorName, tweetedAt *string,
	likesCount, retweetsCount, repliesCount *int, enriched bool) error {
	if b == nil {
		return fmt.Errorf("cannot add Tweet properties because given b is nil")
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

	// Set Tweet-specific properties
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
		// Try to parse the tweetedAt string to ensure it's in RFC3339 format
		parsedTime, err := time.Parse(time.RFC3339, *tweetedAt)
		if err != nil {
			// If it's not already in RFC3339, try other formats
			formats := []string{
				"2006-01-02 15:04:05 -0700 MST",
				"2006-01-02 15:04:05 +0000 UTC",
				"2006-01-02 15:04:05",
				"2006-01-02",
				time.ANSIC,
				time.UnixDate,
				time.RubyDate,
				time.RFC822,
				time.RFC822Z,
				time.RFC850,
				time.RFC1123,
				time.RFC1123Z,
				time.RFC3339Nano,
				time.Kitchen,
				time.Stamp,
				time.StampMilli,
				time.StampMicro,
				time.StampNano,
			}

			var parsedSuccessfully bool
			for _, format := range formats {
				if t, err := time.Parse(format, *tweetedAt); err == nil {
					parsedTime = t
					parsedSuccessfully = true
					break
				}
			}

			if !parsedSuccessfully {
				// If we can't parse the time, just use the string as is
				if err := b.Properties.ReplaceValue(PropertyKeyPublishedAt, *tweetedAt); err != nil {
					return fmt.Errorf("failed to set tweeted at property: %w", err)
				}
			} else {
				// Format the parsed time as RFC3339
				rfc3339Time := parsedTime.Format(time.RFC3339)
				if err := b.Properties.ReplaceValue(PropertyKeyPublishedAt, rfc3339Time); err != nil {
					return fmt.Errorf("failed to set tweeted at property: %w", err)
				}
			}
		} else {
			// The string was already in RFC3339 format, use it directly
			if err := b.Properties.ReplaceValue(PropertyKeyPublishedAt, *tweetedAt); err != nil {
				return fmt.Errorf("failed to set tweeted at property: %w", err)
			}
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

	// Set the enriched flag
	if err := b.Properties.ReplaceValue(PropertyKeyEnriched, enriched); err != nil {
		return fmt.Errorf("failed to set enriched property: %w", err)
	}

	return nil
}

// RenderTweetProperties formats a tweet block's properties as a markdown string
func RenderTweetProperties(b Block) string {
	// Get basic properties
	urlValue, hasURL := b.Properties.Get(PropertyKeyURL)
	url, urlOk := urlValue.(string)

	titleValue, hasTitle := b.Properties.Get(PropertyKeyTitle)
	title, titleOk := titleValue.(string)

	// Get Tweet-specific properties
	usernameValue, hasUsername := b.Properties.Get(PropertyKeyUsername)
	username, usernameOk := usernameValue.(string)

	authorNameValue, hasAuthorName := b.Properties.Get(PropertyKeyAuthorName)
	authorName, authorNameOk := authorNameValue.(string)

	tweetedAtValue, hasTweetedAt := b.Properties.Get(PropertyKeyPublishedAt)
	tweetedAt, tweetedAtOk := tweetedAtValue.(string)

	// Get tweet type flags
	isRetweetValue, hasIsRetweet := b.Properties.Get(PropertyKeyIsRetweet)
	isReplyValue, hasIsReply := b.Properties.Get(PropertyKeyIsReply)
	isQuoteValue, hasIsQuote := b.Properties.Get(PropertyKeyIsQuote)

	// Get source if available
	sourceValue, hasSource := b.Properties.Get(PropertyKeySource)
	source, sourceOk := sourceValue.(string)

	// Get engagement metrics
	likesValue, hasLikes := b.Properties.Get(PropertyKeyLikeCount)
	retweetsValue, hasRetweets := b.Properties.Get(PropertyKeyRetweetCount)
	repliesValue, hasReplies := b.Properties.Get(PropertyKeyCommentCount)
	quotesValue, hasQuotes := b.Properties.Get(PropertyKeyQuoteCount)

	// Get media information
	_, _ = b.Properties.Get(PropertyKeyHasMedia) // We don't need the value directly, just checking presence
	mediaURLsValue, hasMediaURLs := b.Properties.Get(PropertyKeyMediaURLs)
	mediaURLs, mediaURLsOk := mediaURLsValue.([]string)
	if !mediaURLsOk {
		// Try to get it as an array of interfaces
		if mediaURLsArray, ok := b.Properties.GetArray(PropertyKeyMediaURLs); ok {
			mediaURLs = make([]string, 0, len(mediaURLsArray))
			for _, v := range mediaURLsArray {
				if str, ok := v.(string); ok {
					mediaURLs = append(mediaURLs, str)
				}
			}
			mediaURLsOk = true
		}
	}

	// Get external URLs if available
	externalURLsValue, hasExternalURLs := b.Properties.Get(PropertyKeyExternalURLs)
	externalURLs, externalURLsOk := externalURLsValue.([]string)
	if !externalURLsOk {
		// Try to get it as an array of interfaces
		if externalURLsArray, ok := b.Properties.GetArray(PropertyKeyExternalURLs); ok {
			externalURLs = make([]string, 0, len(externalURLsArray))
			for _, v := range externalURLsArray {
				if str, ok := v.(string); ok {
					externalURLs = append(externalURLs, str)
				}
			}
			externalURLsOk = true
		}
	}

	// Format as markdown with Tweet information
	var result string

	// Add tweet type indicators
	var tweetTypeIndicators []string
	if hasIsRetweet {
		if isRetweet, ok := isRetweetValue.(bool); ok && isRetweet {
			tweetTypeIndicators = append(tweetTypeIndicators, "Retweet")
		}
	}
	if hasIsReply {
		if isReply, ok := isReplyValue.(bool); ok && isReply {
			tweetTypeIndicators = append(tweetTypeIndicators, "Reply")
		}
	}
	if hasIsQuote {
		if isQuote, ok := isQuoteValue.(bool); ok && isQuote {
			tweetTypeIndicators = append(tweetTypeIndicators, "Quote Tweet")
		}
	}

	// Add tweet type indicator if any
	if len(tweetTypeIndicators) > 0 {
		typeStr := ""
		for i, indicator := range tweetTypeIndicators {
			if i > 0 {
				typeStr += ", "
			}
			typeStr += indicator
		}
		result += fmt.Sprintf("**%s**\n\n", typeStr)
	}

	// Start with author information
	if hasUsername && usernameOk && username != "" {
		if hasAuthorName && authorNameOk && authorName != "" {
			result += fmt.Sprintf("**%s** (@%s)", authorName, username)
		} else {
			result += fmt.Sprintf("@%s", username)
		}

		// Add tweet timestamp if available
		if hasTweetedAt && tweetedAtOk && tweetedAt != "" {
			result += fmt.Sprintf(" • %s", tweetedAt)
		}

		// Add source if available
		if hasSource && sourceOk && source != "" {
			result += fmt.Sprintf(" • via %s", source)
		}

		result += "\n\n"
	}

	// Add the tweet content
	if hasTitle && titleOk && title != "" {
		result += title + "\n\n"
	}

	// Add media links if available
	if hasMediaURLs && mediaURLsOk && len(mediaURLs) > 0 {
		result += "**Media**:\n"
		for _, mediaURL := range mediaURLs {
			result += fmt.Sprintf("![Media](%s)\n", mediaURL)
		}
		result += "\n"
	}

	// Add external links if available
	if hasExternalURLs && externalURLsOk && len(externalURLs) > 0 {
		result += "**Links**:\n"
		for _, extURL := range externalURLs {
			result += fmt.Sprintf("- [%s](%s)\n", extURL, extURL)
		}
		result += "\n"
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

	if hasQuotes {
		quotes, ok := quotesValue.(int)
		if ok {
			metrics = append(metrics, fmt.Sprintf("🔁 %d", quotes))
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

// GetTweetProperties returns a list of all tweet property keys
func GetTweetProperties() []string {
	return []string{
		PropertyKeyTitle,          // Tweet text
		PropertyKeyURL,            // Tweet URL
		PropertyKeyDescription,    // Additional context/description
		PropertyKeyImageURL,       // Tweet image if any
		PropertyKeyTweetID,        // Unique tweet ID
		PropertyKeyUsername,       // Tweet handle (@username)
		PropertyKeyAuthorID,       // Author ID
		PropertyKeyAuthorName,     // Display name of the author
		PropertyKeyPublishedAt,    // When the tweet was posted
		PropertyKeyLikeCount,      // Number of likes
		PropertyKeyRetweetCount,   // Number of retweets
		PropertyKeyCommentCount,   // Number of replies
		PropertyKeyQuoteCount,     // Number of quotes
		PropertyKeyConversationID, // Conversation ID
		PropertyKeyLanguage,       // Language of the tweet
		PropertyKeySource,         // Client used to post the tweet
		PropertyKeyHasMedia,       // Whether the tweet has media
		PropertyKeyMediaCount,     // Number of media items
		PropertyKeyMediaURLs,      // URLs of media
		PropertyKeyMediaInfo,      // Details about media
		PropertyKeyExternalURLs,   // External URLs
		PropertyKeyIsRetweet,      // Whether it's a retweet
		PropertyKeyIsReply,        // Whether it's a reply
		PropertyKeyIsQuote,        // Whether it's a quote tweet
		PropertyKeyEnriched,       // Whether the tweet has been enriched
	}
}

// ParseTweetDataFromJSON parses a json.RawMessage into a structured.DataTweet struct
func ParseTweetDataFromJSON(rawJSON json.RawMessage) (*structured.DataTweet, error) {
	if len(rawJSON) == 0 {
		return nil, fmt.Errorf("cannot parse tweet data from empty JSON")
	}

	var tweetData structured.DataTweet
	if err := json.Unmarshal(rawJSON, &tweetData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tweet data: %w", err)
	}

	return &tweetData, nil
}

// AddTweetPropertiesFromJSON parses a json.RawMessage into a structured.DataTweet struct
// and adds the tweet properties to the given block
func AddTweetPropertiesFromJSON(block *Block, rawJSON json.RawMessage) error {
	if block == nil {
		return fmt.Errorf("cannot add tweet properties because given block is nil")
	}

	if len(rawJSON) == 0 {
		return fmt.Errorf("cannot add tweet properties because given JSON is empty")
	}

	tweetData, err := ParseTweetDataFromJSON(rawJSON)
	if err != nil {
		return fmt.Errorf("failed to parse tweet data from JSON: %w", err)
	}

	return AddTweetPropertiesFromDataTweet(block, tweetData)
}
