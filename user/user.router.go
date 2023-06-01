package user

import (
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	router.Post("/register", checkFieldRegister, isPasswordSameAsVerifyPassword, IsEmailAlreadyTaken, register)
	router.Post("/login", checkFieldLogin, UserExist, login)
}
