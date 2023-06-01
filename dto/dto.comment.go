package dto

import validation "github.com/go-ozzo/ozzo-validation"

type Comment struct {
	Content string `json:"content"`
	PostId  string `json:"post_id"`
}

func (l Comment) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Content, validation.Required, validation.Length(1, 50)),
		validation.Field(&l.PostId, validation.Required, validation.Length(1, 50)),
	)
}
