package blocks

// PageData defines an interface for objects that can be converted to DataMovie
// This avoids direct dependency on scrappers.Page type while allowing conversion
type PageData interface {
	// GetURL returns the original URL
	GetURL() string

	// GetTitle returns the page title
	GetTitle() string

	// GetDescription returns the short description
	GetDescription() string

	// GetBody returns the full body content
	GetBody() string

	// GetImage returns the image URL
	GetImage() string

	// GetCustomMetadata returns a map of additional metadata
	GetCustomMetadata() map[string]string
}
