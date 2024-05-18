package model

type Comment struct {
	Model
	Author  string `json:"author"`
	Content string `json:"content"`
}
