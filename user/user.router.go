package user

import (
	"backend/dto"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func SetupUserRoutes(router fiber.Router) {
	router.Post("/register", checkFieldRegister, isPasswordSameAsVerifyPassword, IsEmailAlreadyTaken, register)
	router.Post("/login", checkFieldLogin, login)

	router.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(dto.YnoverflowResponse{
				Message: "PLease wait 30 seconds before trying again",
				Code:    fiber.StatusTooManyRequests,
				Data:    nil,
			})
		},
	}))
	router.Put("upload-profile-picture", utils.CheckToken, utils.DecodeToken, uploadProfilePicture)
}
