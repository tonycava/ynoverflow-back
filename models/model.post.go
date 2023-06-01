package models

type Post struct {
	Base
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  string `json:"userId"`
}
