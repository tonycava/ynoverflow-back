package utils

import (
	"backend/dto"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateAccessToken(id, email string) string {
	key, _ := os.LookupEnv("PRIVATE_KEY")
	jwtSecret := []byte(key)

	claims := dto.Claims{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
		},
	}

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tk.SignedString(jwtSecret)
	return token
}
