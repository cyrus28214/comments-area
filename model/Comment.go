package model

type Comment struct {
	Model
	AuthorID uint   `json:"author_id"`
	Content  string `json:"content"`
}
