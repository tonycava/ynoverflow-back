package user

import (
	"backend/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"regexp"
	"time"
)

var jwtKey = []byte(os.Getenv("PRIVATE_KEY"))

func secureAuth() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("access_token")
		claims := new(Claims)

		token, _ := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if token.Valid {
			if claims.ExpiresAt.Unix() < time.Now().Unix() {
				return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
					"error": "Access token expired",
				})
			}
		} else {
			c.ClearCookie("access_token", "refresh_token")
			return c.SendStatus(fiber.StatusForbidden)
		}
		c.Locals("uuid", claims.UUID)
		return c.Next()
	}
}

func ValidateRegister(u *User) *UserError {
	e := UserError{}
	e.Err, e.Username = utils.IsEmpty(u.Username)

	if !govalidator.IsEmail(u.Email) {
		e.Err, e.Email = true, "Invalid email"
	}

	re := regexp.MustCompile("\\d")
	if !(len(u.Password) >= 8 && govalidator.HasLowerCase(u.Password) && govalidator.HasUpperCase(u.Password) && re.MatchString(u.Password)) {
		e.Err, e.Password = true, "Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter and one number"
	}

	return &e
}
