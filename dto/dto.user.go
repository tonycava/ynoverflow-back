package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Register struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verifyPassword"`
}

func (l Register) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required, validation.Length(1, 50), is.Email),
		validation.Field(&l.Username, validation.Required, validation.Length(1, 50)),
		validation.Field(&l.Password, validation.Required, validation.Length(1, 50)),
		validation.Field(&l.VerifyPassword, validation.Required, validation.Length(1, 50)),
	)
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required, validation.Length(1, 50), is.Email),
		validation.Field(&l.Password, validation.Required, validation.Length(1, 50)),
	)
}
