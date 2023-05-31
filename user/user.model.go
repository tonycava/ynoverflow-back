package user

import (
	"backend/models"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	models.Base
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}

type UserError struct {
	Err      bool   `json:"error"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.RegisteredClaims
	UUID string `json:"uuid" gorm:"primaryKey"`
}
