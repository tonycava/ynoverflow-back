package models

type User struct {
	Base
	Email    string    `gorm:"unique" json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Posts    []Post    `json:"posts"`
	Comments []Comment `json:"comments"`
	View     []View    `json:"views"`
}
