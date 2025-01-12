package models

type Author struct {
	AuthorID  int    `json:"author_id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}
