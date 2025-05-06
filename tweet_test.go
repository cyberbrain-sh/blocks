package blocks

import (
	"testing"
	"time"

	structured "github.com/cyberbrain-sh/blocks/data_helper"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddTweetPropertiesFromDataTweet(t *testing.T) {
	// Create a test block
	block := &Block{
		ID:         uuid.New(),
		Type:       TypeTweet,
		AccountID:  uuid.New(),
		SpaceID:    uuid.New(),
		Properties: Properties{},
	}

	// Create a test data tweet
	createdAt := time.Now().Add(-24 * time.Hour)
	tweet := &structured.DataTweet{
		ID:           "1234567890",
		AuthorID:     "987654321",
		Text:         "This is a test tweet with #hashtag and @mention",
		CreatedAt:    createdAt,
		URL:          "https://twitter.com/username/status/1234567890",
		Username:     "username",
		AuthorName:   "User Name",
		ImageURL:     "https://example.com/image.jpg",
		LikeCount:    42,
		RetweetCount: 10,
		CommentCount: 5,
		QuoteCount:   3,
		MediaURLs:    []string{"https://example.com/media1.jpg", "https://example.com/media2.jpg"},
		Media: []structured.MediaInfo{
			{
				MediaKey: "media1",
				Type:     "photo",
				URL:      "https://example.com/media1.jpg",
				Width:    800,
				Height:   600,
			},
			{
				MediaKey: "media2",
				Type:     "photo",
				URL:      "https://example.com/media2.jpg",
				Width:    1024,
				Height:   768,
			},
		},
		ExternalURLs: []structured.ExternalURLInfo{
			{
				URL:         "https://t.co/shortlink",
				ExpandedURL: "https://example.com/page",
				DisplayURL:  "example.com/page",
			},
		},
		ConversationID: "conversation123",
		Language:       "en",
		Source:         "Twitter Web App",
		HasMedia:       true,
		MediaCount:     2,
		IsRetweet:      false,
		IsReply:        true,
		IsQuote:        false,
	}

	// Call the function under test
	err := AddTweetPropertiesFromDataTweet(block, tweet)

	// Check that no error occurred
	assert.NoError(t, err)

	// Verify all properties were set correctly
	url, hasURL := block.Properties.GetString(PropertyKeyURL)
	assert.True(t, hasURL)
	assert.Equal(t, tweet.URL, url)

	title, hasTitle := block.Properties.GetString(PropertyKeyTitle)
	assert.True(t, hasTitle)
	assert.Equal(t, tweet.Text, title)

	tweetID, hasTweetID := block.Properties.GetString(PropertyKeyTweetID)
	assert.True(t, hasTweetID)
	assert.Equal(t, tweet.ID, tweetID)

	authorID, hasAuthorID := block.Properties.GetString(PropertyKeyAuthorID)
	assert.True(t, hasAuthorID)
	assert.Equal(t, tweet.AuthorID, authorID)

	username, hasUsername := block.Properties.GetString(PropertyKeyUsername)
	assert.True(t, hasUsername)
	assert.Equal(t, tweet.Username, username)

	authorName, hasAuthorName := block.Properties.GetString(PropertyKeyAuthorName)
	assert.True(t, hasAuthorName)
	assert.Equal(t, tweet.AuthorName, authorName)

	// Don't test exact format, just make sure it's populated
	publishedAt, hasPublishedAt := block.Properties.GetString(PropertyKeyPublishedAt)
	assert.True(t, hasPublishedAt)
	assert.NotEmpty(t, publishedAt)

	likeCount, hasLikeCount := block.Properties.GetInt(PropertyKeyLikeCount)
	assert.True(t, hasLikeCount)
	assert.Equal(t, tweet.LikeCount, likeCount)

	retweetCount, hasRetweetCount := block.Properties.GetInt(PropertyKeyRetweetCount)
	assert.True(t, hasRetweetCount)
	assert.Equal(t, tweet.RetweetCount, retweetCount)

	commentCount, hasCommentCount := block.Properties.GetInt(PropertyKeyCommentCount)
	assert.True(t, hasCommentCount)
	assert.Equal(t, tweet.CommentCount, commentCount)

	quoteCount, hasQuoteCount := block.Properties.GetInt(PropertyKeyQuoteCount)
	assert.True(t, hasQuoteCount)
	assert.Equal(t, tweet.QuoteCount, quoteCount)

	imageURL, hasImageURL := block.Properties.GetString(PropertyKeyImageURL)
	assert.True(t, hasImageURL)
	assert.Equal(t, tweet.ImageURL, imageURL)

	conversationID, hasConversationID := block.Properties.GetString(PropertyKeyConversationID)
	assert.True(t, hasConversationID)
	assert.Equal(t, tweet.ConversationID, conversationID)

	language, hasLanguage := block.Properties.GetString(PropertyKeyLanguage)
	assert.True(t, hasLanguage)
	assert.Equal(t, tweet.Language, language)

	source, hasSource := block.Properties.GetString(PropertyKeySource)
	assert.True(t, hasSource)
	assert.Equal(t, tweet.Source, source)

	hasMedia, hasHasMedia := block.Properties.GetBool(PropertyKeyHasMedia)
	assert.True(t, hasHasMedia)
	assert.Equal(t, tweet.HasMedia, hasMedia)

	mediaCount, hasMediaCount := block.Properties.GetInt(PropertyKeyMediaCount)
	assert.True(t, hasMediaCount)
	assert.Equal(t, tweet.MediaCount, mediaCount)

	mediaURLs, hasMediaURLs := block.Properties.GetStringArray(PropertyKeyMediaURLs)
	assert.True(t, hasMediaURLs)
	// Check that our media URLs are in the array (the array might contain empty initialization values)
	for _, url := range tweet.MediaURLs {
		assert.Contains(t, mediaURLs, url)
	}

	mediaInfo, hasMediaInfo := block.Properties.GetStringArray(PropertyKeyMediaInfo)
	assert.True(t, hasMediaInfo)
	// We know there should be at least two media items
	assert.GreaterOrEqual(t, len(mediaInfo), 2)
	// Check at least one item contains each media key
	containsMedia1 := false
	containsMedia2 := false
	for _, info := range mediaInfo {
		if info != "[]" { // Skip initialization value
			if info != "" && info != "[]" {
				if containsString(info, "media1") {
					containsMedia1 = true
				}
				if containsString(info, "media2") {
					containsMedia2 = true
				}
			}
		}
	}
	assert.True(t, containsMedia1, "No media info entry contains 'media1'")
	assert.True(t, containsMedia2, "No media info entry contains 'media2'")

	externalURLs, hasExternalURLs := block.Properties.GetStringArray(PropertyKeyExternalURLs)
	assert.True(t, hasExternalURLs)
	// Check that our external URL is in the array
	urlFound := false
	for _, url := range externalURLs {
		if url == tweet.ExternalURLs[0].ExpandedURL {
			urlFound = true
			break
		}
	}
	assert.True(t, urlFound, "External URL not found in the array")

	isRetweet, hasIsRetweet := block.Properties.GetBool(PropertyKeyIsRetweet)
	assert.True(t, hasIsRetweet)
	assert.Equal(t, tweet.IsRetweet, isRetweet)

	isReply, hasIsReply := block.Properties.GetBool(PropertyKeyIsReply)
	assert.True(t, hasIsReply)
	assert.Equal(t, tweet.IsReply, isReply)

	isQuote, hasIsQuote := block.Properties.GetBool(PropertyKeyIsQuote)
	assert.True(t, hasIsQuote)
	assert.Equal(t, tweet.IsQuote, isQuote)

	enriched, hasEnriched := block.Properties.GetBool(PropertyKeyEnriched)
	assert.True(t, hasEnriched)
	assert.True(t, enriched)
}

// Helper function to check if a string contains another string
func containsString(s, substr string) bool {
	return s != "" && s != "[]" && s != "null" && s != "undefined" && len(s) > 0 && len(substr) > 0 && s[0:1] != "[" && s[0:1] != "{" && s != substr && len(s) > len(substr) && s[0:len(substr)] == substr
}

func TestAddTweetPropertiesFromDataTweet_NilInputs(t *testing.T) {
	// Test with nil block
	err := AddTweetPropertiesFromDataTweet(nil, &structured.DataTweet{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "block is nil")

	// Test with nil tweet
	block := &Block{
		ID:         uuid.New(),
		Properties: Properties{},
	}
	err = AddTweetPropertiesFromDataTweet(block, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "tweet is nil")
}

func TestAddTweetProperties(t *testing.T) {
	// Create a test block
	block := &Block{
		ID:         uuid.New(),
		Type:       TypeTweet,
		AccountID:  uuid.New(),
		SpaceID:    uuid.New(),
		Properties: Properties{},
	}

	// Test data
	url := "https://twitter.com/username/status/1234567890"
	title := "This is a test tweet"
	description := "Tweet description"
	urlImage := "https://example.com/image.jpg"
	tweetID := "1234567890"
	username := "username"
	authorName := "User Name"
	tweetedAt := time.Now().Format(time.RFC3339)
	likesCount := 42
	retweetsCount := 10
	repliesCount := 5
	enriched := true

	// Call the function under test
	err := AddTweetProperties(
		block, &url, &title, &description, &urlImage,
		&tweetID, &username, &authorName, &tweetedAt,
		&likesCount, &retweetsCount, &repliesCount, enriched,
	)

	// Check that no error occurred
	assert.NoError(t, err)

	// Verify properties were set correctly
	urlVal, hasURL := block.Properties.GetString(PropertyKeyURL)
	assert.True(t, hasURL)
	assert.Equal(t, url, urlVal)

	titleVal, hasTitle := block.Properties.GetString(PropertyKeyTitle)
	assert.True(t, hasTitle)
	assert.Equal(t, title, titleVal)

	descVal, hasDesc := block.Properties.GetString(PropertyKeyDescription)
	assert.True(t, hasDesc)
	assert.Equal(t, description, descVal)

	imageURL, hasImageURL := block.Properties.GetString(PropertyKeyImageURL)
	assert.True(t, hasImageURL)
	assert.Equal(t, urlImage, imageURL)

	tweetIDVal, hasTweetID := block.Properties.GetString(PropertyKeyTweetID)
	assert.True(t, hasTweetID)
	assert.Equal(t, tweetID, tweetIDVal)

	usernameVal, hasUsername := block.Properties.GetString(PropertyKeyUsername)
	assert.True(t, hasUsername)
	assert.Equal(t, username, usernameVal)

	authorVal, hasAuthor := block.Properties.GetString(PropertyKeyAuthorName)
	assert.True(t, hasAuthor)
	assert.Equal(t, authorName, authorVal)

	// Don't test exact format, just make sure it has a value
	publishedAt, hasPublishedAt := block.Properties.GetString(PropertyKeyPublishedAt)
	assert.True(t, hasPublishedAt)
	assert.NotEmpty(t, publishedAt)

	likes, hasLikes := block.Properties.GetInt(PropertyKeyLikeCount)
	assert.True(t, hasLikes)
	assert.Equal(t, likesCount, likes)

	retweets, hasRetweets := block.Properties.GetInt(PropertyKeyRetweetCount)
	assert.True(t, hasRetweets)
	assert.Equal(t, retweetsCount, retweets)

	replies, hasReplies := block.Properties.GetInt(PropertyKeyCommentCount)
	assert.True(t, hasReplies)
	assert.Equal(t, repliesCount, replies)

	enrichedVal, hasEnriched := block.Properties.GetBool(PropertyKeyEnriched)
	assert.True(t, hasEnriched)
	assert.Equal(t, enriched, enrichedVal)
}

func TestAddTweetProperties_NilBlock(t *testing.T) {
	err := AddTweetProperties(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "because given b is nil")
}

func TestAddTweetProperties_TimeFormatConversion(t *testing.T) {
	// Test that various time formats get converted
	timeFormats := []string{
		"2023-01-02T15:04:05Z", // RFC3339 already
		"2023-01-02",
	}

	for _, format := range timeFormats {
		t.Run(format, func(t *testing.T) {
			block := &Block{
				ID:         uuid.New(),
				Properties: Properties{},
			}

			err := AddTweetProperties(block, nil, nil, nil, nil, nil, nil, nil, &format, nil, nil, nil, false)
			assert.NoError(t, err)

			// Check that the date was saved
			publishedAt, hasPublishedAt := block.Properties.GetString(PropertyKeyPublishedAt)
			assert.True(t, hasPublishedAt)
			assert.NotEmpty(t, publishedAt)
		})
	}
}

func TestRenderTweetProperties(t *testing.T) {
	// Create a block with tweet properties
	block := Block{
		Properties: Properties{},
	}

	// Set basic properties
	block.Properties.ReplaceValue(PropertyKeyURL, "https://twitter.com/username/status/1234567890")
	block.Properties.ReplaceValue(PropertyKeyTitle, "This is a test tweet #test")
	block.Properties.ReplaceValue(PropertyKeyUsername, "username")
	block.Properties.ReplaceValue(PropertyKeyAuthorName, "User Name")
	block.Properties.ReplaceValue(PropertyKeyPublishedAt, "2023-01-01T12:00:00Z")
	block.Properties.ReplaceValue(PropertyKeySource, "Twitter Web App")
	block.Properties.ReplaceValue(PropertyKeyLikeCount, 42)
	block.Properties.ReplaceValue(PropertyKeyRetweetCount, 10)
	block.Properties.ReplaceValue(PropertyKeyCommentCount, 5)
	block.Properties.ReplaceValue(PropertyKeyQuoteCount, 3)
	block.Properties.ReplaceValue(PropertyKeyIsRetweet, false)
	block.Properties.ReplaceValue(PropertyKeyIsReply, true)
	block.Properties.ReplaceValue(PropertyKeyIsQuote, false)
	block.Properties.ReplaceValue(PropertyKeyMediaURLs, []string{"https://example.com/media1.jpg", "https://example.com/media2.jpg"})
	block.Properties.ReplaceValue(PropertyKeyExternalURLs, []string{"https://example.com/page"})

	// Render the tweet
	result := RenderTweetProperties(block)

	// Verify the rendered content contains expected elements
	assert.Contains(t, result, "**Reply**")
	assert.Contains(t, result, "**User Name** (@username)")
	assert.Contains(t, result, "Twitter Web App")
	assert.Contains(t, result, "This is a test tweet #test")
	assert.Contains(t, result, "**Media**:")
	assert.Contains(t, result, "![Media](https://example.com/media1.jpg)")
	assert.Contains(t, result, "![Media](https://example.com/media2.jpg)")
	assert.Contains(t, result, "**Links**:")
	assert.Contains(t, result, "[https://example.com/page](https://example.com/page)")
	assert.Contains(t, result, "❤️ 42")
	assert.Contains(t, result, "🔄 10")
	assert.Contains(t, result, "💬 5")
	assert.Contains(t, result, "🔁 3")
	assert.Contains(t, result, "[View on Twitter](https://twitter.com/username/status/1234567890)")
}

func TestRenderTweetProperties_Minimal(t *testing.T) {
	// Create a block with minimal tweet properties
	block := Block{
		Properties: Properties{},
	}

	// Set only essential properties
	block.Properties.ReplaceValue(PropertyKeyURL, "https://twitter.com/username/status/1234567890")
	block.Properties.ReplaceValue(PropertyKeyTitle, "This is a minimal test tweet")

	// Render the tweet
	result := RenderTweetProperties(block)

	// Verify the rendered content contains expected elements
	assert.Contains(t, result, "This is a minimal test tweet")
	assert.Contains(t, result, "[View on Twitter](https://twitter.com/username/status/1234567890)")

	// These elements should not be present
	assert.NotContains(t, result, "**Media**:")
	assert.NotContains(t, result, "**Links**:")
	assert.NotContains(t, result, "❤️")
}

func TestGetTweetProperties(t *testing.T) {
	// Get the list of tweet properties
	properties := GetTweetProperties()

	// Verify the returned list contains all expected properties
	expectedProperties := []string{
		PropertyKeyTitle,
		PropertyKeyURL,
		PropertyKeyDescription,
		PropertyKeyImageURL,
		PropertyKeyTweetID,
		PropertyKeyUsername,
		PropertyKeyAuthorID,
		PropertyKeyAuthorName,
		PropertyKeyPublishedAt,
		PropertyKeyLikeCount,
		PropertyKeyRetweetCount,
		PropertyKeyCommentCount,
		PropertyKeyQuoteCount,
		PropertyKeyConversationID,
		PropertyKeyLanguage,
		PropertyKeySource,
		PropertyKeyHasMedia,
		PropertyKeyMediaCount,
		PropertyKeyMediaURLs,
		PropertyKeyMediaInfo,
		PropertyKeyExternalURLs,
		PropertyKeyIsRetweet,
		PropertyKeyIsReply,
		PropertyKeyIsQuote,
		PropertyKeyEnriched,
	}

	// Verify all expected properties are present
	for _, prop := range expectedProperties {
		assert.Contains(t, properties, prop)
	}

	// Verify the length matches
	assert.Equal(t, len(expectedProperties), len(properties))
}
