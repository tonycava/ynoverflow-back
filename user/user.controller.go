package user

import (
	"backend/dto"
	"backend/models"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func register(c *fiber.Ctx) error {
	var user = c.Locals("user").(dto.Register)

	createdUser := models.User{
		Base:     models.Base{},
		Email:    user.Email,
		Username: user.Username,
		Password: utils.HashPassword(user.Password),
	}
	createdUser = CreateUser(createdUser)

	return c.Status(fiber.StatusOK).JSON(dto.YnoverflowResponse{
		Message: "Successfully registered",
		Code:    fiber.StatusOK,
		Data: fiber.Map{
			"token": utils.GenerateAccessToken(createdUser.ID, createdUser.Email),
		},
	})
}

func login(c *fiber.Ctx) error {
	var user = c.Locals("user").(dto.Login)
	var userFromDB = GetUserByEmail(user.Email)

	if !utils.CheckPasswordHash(user.Password, userFromDB.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: "Invalid credentials",
			Code:    fiber.StatusBadRequest,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.YnoverflowResponse{
		Message: "Successfully logged in",
		Code:    fiber.StatusOK,
		Data: fiber.Map{
			"token": utils.GenerateAccessToken(userFromDB.ID, userFromDB.Email),
		},
	})
}
