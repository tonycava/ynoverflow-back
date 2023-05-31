package router

import (
	"backend/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	user.SetupUserRoutes(app.Group("/user"))
}
