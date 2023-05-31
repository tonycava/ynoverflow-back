package utils

import (
	"backend/database"
	"backend/models"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateTokens() (string, string) {
	claim, accessToken := GenerateAccessClaims()
	refreshToken := GenerateRefreshClaims(claim)

	return accessToken, refreshToken
}

func GenerateAccessClaims() (*models.Claims, string) {

	key, _ := os.LookupEnv("PRIVATE_KEY")

	jwtKey := []byte(key)

	token, _ := jwt.ParseWithClaims(os.Getenv("ACCESS_TOKEN"), &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, token.Raw
	} else {
		return nil, ""
	}
}

func GenerateRefreshClaims(claim *models.Claims) string {

	key, _ := os.LookupEnv("PRIVATE_KEY")
	jwtKey := []byte(key)

	result := database.DBInstance{}.Db.Where("expires_at < ?", time.Now()).Delete(&models.Claims{})

	if result.RowsAffected > 3 {
		result.Delete(&models.Claims{})
	}
	refreshClaim := &models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		},
	}

	database.DBInstance{}.Db.Create(&refreshClaim)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, _ := refreshToken.SignedString(jwtKey)

	return token
}
