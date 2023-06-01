package router

import (
	"backend/comment"
	"backend/post"
	"backend/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	user.SetupUserRoutes(app.Group("/user"))
	post.SetupPostRoutes(app.Group("/post"))
	comment.SetupRoutesComments(app.Group("/comment"))
}
