package utils

import (
	"backend/database"
	"backend/user"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("PRIVATE_KEY"))

func GenerateTokens() (string, string) {
	claim, accessToken := GenerateAccessClaims()
	refreshToken := GenerateRefreshClaims(claim)

	return accessToken, refreshToken
}

func GenerateAccessClaims() (*user.Claims, string) {
	token, _ := jwt.ParseWithClaims(os.Getenv("ACCESS_TOKEN"), &user.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(*user.Claims); ok && token.Valid {
		return claims, token.Raw
	} else {
		return nil, ""
	}
}

func GenerateRefreshClaims(claim *user.Claims) string {
	result := database.DB.Where("expires_at < ?", time.Now()).Delete(&user.Claims{}).Find(&user.Claims{})

	if result.RowsAffected > 3 {
		result.Delete(&user.Claims{})
	}
	refreshClaim := &user.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		},
	}

	database.DB.Create(&refreshClaim)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, _ := refreshToken.SignedString(jwtKey)

	return token
}
