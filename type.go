package blocks

type DataType string

func (d DataType) String() string {
	return string(d)
}

// Types copied from structured package
const (
	// TypeFragment this is initial state block is created at, then it must transition to one of the types
	TypeFragment  DataType = "fragment"
	TypePage      DataType = "page"
	TypeImage     DataType = "image"
	TypeVideo     DataType = "video"
	TypeAudio     DataType = "audio"
	TypeFile      DataType = "file"
	TypeDatabase  DataType = "database"
	TypeLink      DataType = "link"
	TypeToDo      DataType = "to_do"
	TypeEmail     DataType = "email"
	TypeMovie     DataType = "movie"
	TypeInstagram DataType = "instagram"
	TypeYouTube   DataType = "youtube"
	TypeTweet     DataType = "tweet"
	TypeSeries    DataType = "series"
	TypeBook      DataType = "book"
	TypeLine      DataType = "line"
	TypePerson    DataType = "person"
	TypePlace     DataType = "place"
	// add all text blocks as well once we are ready
	TypeParagraph        DataType = "paragraph"
	TypeHeader1          DataType = "heading_1"
	TypeHeader2          DataType = "heading_2"
	TypeHeader3          DataType = "heading_3"
	TypeHeader4          DataType = "heading_4"
	TypeHeader5          DataType = "heading_5"
	TypeHeader6          DataType = "heading_6"
	TypeBulletListItem   DataType = "bullet_list_item"
	TypeNumberedListItem DataType = "numbered_list_item"
)

// IsValid checks if the DataType is one of the defined valid types
func (d DataType) IsValid() bool {
	switch d {
	case TypeFragment,
		TypePage,
		TypeImage,
		TypeVideo,
		TypeAudio,
		TypeFile,
		TypeDatabase,
		TypeLink,
		TypeToDo,
		TypeEmail,
		TypeMovie,
		TypeInstagram,
		TypeYouTube,
		TypeTweet,
		TypeSeries,
		TypeBook,
		TypeLine,
		TypePerson,
		TypePlace,
		TypeParagraph,
		TypeHeader1,
		TypeHeader2,
		TypeHeader3,
		TypeHeader4,
		TypeHeader5,
		TypeHeader6,
		TypeBulletListItem,
		TypeNumberedListItem:
		return true
	}
	return false
}

type BlockContentType string

func (b BlockContentType) String() string {
	return string(b)
}

const BlockContentTypeTextual BlockContentType = "textual"
const BlockContentTypeStructural BlockContentType = "structural"

func (d DataType) ContentType() BlockContentType {
	switch d {
	case TypePage, TypeParagraph, TypeHeader1, TypeHeader2, TypeHeader3, TypeHeader4, TypeHeader5, TypeHeader6, TypeBulletListItem, TypeNumberedListItem, TypeLine:
		return BlockContentTypeTextual
	}
	return BlockContentTypeStructural
}

func (d DataType) IsFinal() bool {
	switch d {
	case TypeMovie, TypeSeries, TypeBook, TypePlace, TypePerson, TypeImage, TypeVideo, TypeAudio, TypeFile:
		return true
	default:
		return false
	}
}
