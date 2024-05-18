package model

type Comment struct {
	Model
	Author  uint   `json:"author"`
	Content string `json:"content"`
}
