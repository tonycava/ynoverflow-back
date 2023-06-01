package models

type Comment struct {
	Base
	Content string `json:"content"`
	PostId  string `json:"posId"`
	UserId  string `json:"userId"`
}
