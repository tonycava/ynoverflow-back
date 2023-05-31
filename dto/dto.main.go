package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/golang-jwt/jwt/v5"
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

type YnoverflowResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type Claims struct {
	jwt.RegisteredClaims
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	Email     string `json:"email"`
}
