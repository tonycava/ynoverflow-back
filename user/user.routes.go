package user

import (
	"backend/database"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

var USER fiber.Router

func SetupUserRoutes() {
	USER.Post("/register", CreateUser)
}

func CreateUser(c *fiber.Ctx) error {
	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return c.JSON(fiber.Map{
			"error": "Please review your input",
		})
	}

	errors := ValidateRegister(u)
	if errors.Err {
		return c.JSON(errors)
	}

	if count := database.DB.Where(User{
		Email: u.Email,
	}).First(new(User)).RowsAffected; count > 0 {
		errors.Err, errors.Email = true, "Email already exists"
	}

	if count := database.DB.Where(User{
		Username: u.Username,
	}).First(new(User)).RowsAffected; count > 0 {
		errors.Err, errors.Username = true, "Username already exists"
	}

	if errors.Err {
		return c.JSON(errors)
	}

	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, rand.Intn(10)+10)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	u.Password = string(hashedPassword)
	if err := database.DB.Create(&u).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	accessToken, refreshToken := utils.GenerateTokens()
	accessCookie, refreshCookie := utils.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
