package user

import (
	"backend/dto"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func IsEmailAlreadyTaken(c *fiber.Ctx) error {
	userData := c.Locals("user").(dto.Register)
	var user = GetUserByEmail(userData.Email)
	if user.Email != "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: "Email already taken",
			Code:    fiber.StatusBadRequest,
			Data:    nil,
		})
	}
	return c.Next()
}

func isPasswordSameAsVerifyPassword(c *fiber.Ctx) error {
	userData := c.Locals("user").(dto.Register)
	if userData.Password != userData.VerifyPassword {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: "Password and verify password are not the same",
			Code:    fiber.StatusBadRequest,
			Data:    nil,
		})
	}
	return c.Next()
}

func checkFieldRegister(c *fiber.Ctx) error {
	var register dto.Register
	var err = c.BodyParser(&register)
	var errorValidation = register.Validate()

	if err != nil || errorValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: utils.Parse(errorValidation.Error()),
			Code:    fiber.StatusBadRequest,
			Data:    nil,
		})
	}

	c.Locals("user", register)
	return c.Next()
}

func UserExist(c *fiber.Ctx) error {
	var user = c.Locals("user").(dto.Login)
	var userFromDB = GetUserByEmail(user.Email)
	if userFromDB.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: "User does not exist",
			Code:    fiber.StatusBadRequest,
			Data:    nil,
		})
	}
	return c.Next()
}

func checkFieldLogin(c *fiber.Ctx) error {
	var login dto.Login
	var err = c.BodyParser(&login)
	var errorValidation = login.Validate()

	if err != nil || errorValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: utils.Parse(errorValidation.Error()),
			Code:    fiber.StatusBadRequest,
			Data:    nil,
		})
	}

	c.Locals("user", login)
	return c.Next()
}
