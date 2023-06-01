package utils

import (
	"backend/dto"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func DecodeToken(c *fiber.Ctx) error {
	tokenHeader := c.Locals("token").(*jwt.Token)
	jwtSecret, _ := os.LookupEnv("JWT_SECRET")

	var AccessToken map[string]string
	stringify, _ := json.Marshal(&tokenHeader)
	_ = json.Unmarshal(stringify, &AccessToken)

	token, _ := jwt.ParseWithClaims(AccessToken["Raw"], &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	decodedToken := token.Claims.(*dto.Claims)

	c.Locals("decodedToken", decodedToken)

	return c.Next()
}

func CheckToken(c *fiber.Ctx) error {
	tokenString := c.GetReqHeaders()["Authorization"]

	jwtSecret, _ := os.LookupEnv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.YnoverflowResponse{
			Message: "Wrong token",
			Code:    fiber.StatusUnauthorized,
		})
	}

	if _, ok := token.Claims.(*dto.Claims); ok && token.Valid {
		c.Locals("token", token)
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.YnoverflowResponse{
			Message: "Wrong token",
			Code:    fiber.StatusUnauthorized,
		})
	}
}

func GenerateAccessToken(id, email string) string {
	key, _ := os.LookupEnv("JWT_SECRET")
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
