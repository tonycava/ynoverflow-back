package utils

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetAuthCookies(accessToken, refreshToken string) (*fiber.Cookie, *fiber.Cookie) {
	accessCookie := &fiber.Cookie{
		Name:    "access_token",
		Value:   accessToken,
		Expires: time.Now().Add(24 * time.Hour),
		Secure:  true,
	}

	refreshCookie := &fiber.Cookie{
		Name:    "refresh_token",
		Value:   refreshToken,
		Expires: time.Now().Add(10 * 24 * time.Hour),
		Secure:  true,
	}

	return accessCookie, refreshCookie
}
