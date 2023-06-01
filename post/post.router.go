package post

import (
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(router fiber.Router) {
	router.Get("/", GetAllPostsController)
	router.Get("/:userId", CheckUserIdParams, GetPostsByUserIdController)
	router.Post("/", utils.CheckToken, utils.DecodeToken, CheckFieldCreatePost, CreatePostController)
}
