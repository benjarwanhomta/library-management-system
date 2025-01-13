package models

type Authors struct {
	AuthorID  int    `json:"author_id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}
