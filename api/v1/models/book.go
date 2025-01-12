package models

// Books represents books in the library
type Books struct {
	BookID       int    `json:"book_id"`
	Title        string `json:"title"`
	AuthorID     int    `json:"author_id"`
	CategoryID   int    `json:"category_id"`
	PublishYear  int    `json:"publish_year"`
	ISBN         string `json:"isbn"`
	Description  string `json:"description"`
	AvailableQTY int    `json:"available_qty"`

	Author   Author   `gorm:"foreignKey:AuthorID;references:AuthorID"`
	Category Category `gorm:"foreignKey:CategoryID;references:CategoryID"`
}
