package structured

import (
	"time"
)

// MediaInfo represents media information for a tweet
type MediaInfo struct {
	MediaKey   string `json:"media_key"`
	Type       string `json:"type"`
	URL        string `json:"url,omitempty"`
	PreviewURL string `json:"preview_image_url,omitempty"`
	Width      int    `json:"width,omitempty"`
	Height     int    `json:"height,omitempty"`
	AltText    string `json:"alt_text,omitempty"`
}

// ExternalURLInfo represents an external URL with metadata
type ExternalURLInfo struct {
	URL         string `json:"url"`                   // Original short URL as it appears in the tweet
	ExpandedURL string `json:"expanded_url"`          // Expanded/unshortened URL
	DisplayURL  string `json:"display_url"`           // User-friendly display version
	Title       string `json:"title,omitempty"`       // Page title if available
	Description string `json:"description,omitempty"` // Page description if available
	UnwoundURL  string `json:"unwound_url,omitempty"` // Fully unwound URL after all redirects
}

// DataTweet represents structured data for a tweet
type DataTweet struct {
	ID             string            `json:"id"`
	AuthorID       string            `json:"author_id"`
	Text           string            `json:"text"`
	CreatedAt      time.Time         `json:"created_at"`
	URL            string            `json:"url"`
	Username       string            `json:"username,omitempty"`
	AuthorName     string            `json:"author_name,omitempty"`
	ImageURL       string            `json:"image_url,omitempty"`
	MediaURLs      []string          `json:"media_urls,omitempty"`    // All media URLs from the tweet
	Media          []MediaInfo       `json:"media,omitempty"`         // All media objects associated with this tweet
	ExternalURLs   []ExternalURLInfo `json:"external_urls,omitempty"` // External URLs found in the tweet
	LikeCount      int               `json:"like_count,omitempty"`
	RetweetCount   int               `json:"retweet_count,omitempty"`
	CommentCount   int               `json:"comment_count,omitempty"`
	QuoteCount     int               `json:"quote_count,omitempty"`
	ConversationID string            `json:"conversation_id,omitempty"` // ID of the conversation this tweet is part of
	Language       string            `json:"language,omitempty"`        // Language of the tweet
	Source         string            `json:"source,omitempty"`          // Client used to post the tweet
	HasMedia       bool              `json:"has_media"`                 // Whether the tweet has any media
	MediaCount     int               `json:"media_count,omitempty"`     // Number of media items
	IsRetweet      bool              `json:"is_retweet"`                // Whether this is a retweet
	IsReply        bool              `json:"is_reply"`                  // Whether this is a reply
	IsQuote        bool              `json:"is_quote"`                  // Whether this is a quote tweet
}

// NewDataTweet creates a new DataTweet instance with basic fields
func NewDataTweet(id, authorID, text string, createdAt time.Time, url string) *DataTweet {
	return &DataTweet{
		ID:        id,
		AuthorID:  authorID,
		Text:      text,
		CreatedAt: createdAt,
		URL:       url,
	}
}

// NewEnhancedDataTweet creates a new DataTweet instance with all available fields
func NewEnhancedDataTweet(
	id, authorID, text string,
	createdAt time.Time,
	url, username, authorName, imageURL string,
	likeCount, retweetCount, commentCount, quoteCount int,
) *DataTweet {
	return &DataTweet{
		ID:           id,
		AuthorID:     authorID,
		Text:         text,
		CreatedAt:    createdAt,
		URL:          url,
		Username:     username,
		AuthorName:   authorName,
		ImageURL:     imageURL,
		LikeCount:    likeCount,
		RetweetCount: retweetCount,
		CommentCount: commentCount,
		QuoteCount:   quoteCount,
	}
}

// NewCompleteDataTweet creates a comprehensive DataTweet with all fields including media and URLs
func NewCompleteDataTweet(
	id, authorID, text string,
	createdAt time.Time,
	url, username, authorName, imageURL string,
	mediaURLs []string,
	media []MediaInfo,
	externalURLs []ExternalURLInfo,
	likeCount, retweetCount, commentCount, quoteCount int,
	conversationID, language, source string,
	hasMedia bool,
	mediaCount int,
	isRetweet, isReply, isQuote bool,
) *DataTweet {
	return &DataTweet{
		ID:             id,
		AuthorID:       authorID,
		Text:           text,
		CreatedAt:      createdAt,
		URL:            url,
		Username:       username,
		AuthorName:     authorName,
		ImageURL:       imageURL,
		MediaURLs:      mediaURLs,
		Media:          media,
		ExternalURLs:   externalURLs,
		LikeCount:      likeCount,
		RetweetCount:   retweetCount,
		CommentCount:   commentCount,
		QuoteCount:     quoteCount,
		ConversationID: conversationID,
		Language:       language,
		Source:         source,
		HasMedia:       hasMedia,
		MediaCount:     mediaCount,
		IsRetweet:      isRetweet,
		IsReply:        isReply,
		IsQuote:        isQuote,
	}
}
