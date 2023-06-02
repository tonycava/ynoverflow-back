package models

type Post struct {
	Base
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserId   string    `json:"user_id"`
	Comments []Comment `json:"comments"`
	View     []View    `json:"views"`
}
