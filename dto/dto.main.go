package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

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
