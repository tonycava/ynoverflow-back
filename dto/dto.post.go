package dto

import validation "github.com/go-ozzo/ozzo-validation"

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (l Post) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Title, validation.Required, validation.Length(1, 50)),
		validation.Field(&l.Content, validation.Required, validation.Length(1, 50)),
	)
}

type PostDTO struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}
