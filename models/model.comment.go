package models

type Comment struct {
	Base
	Content string `json:"content"`
	PostId  string `json:"post_id"`
	UserId  string `json:"user_id"`
}
