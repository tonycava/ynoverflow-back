package post

import (
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(router fiber.Router) {
	router.Get("/", GetAllPostsController)
	router.Post("/", utils.CheckToken, utils.DecodeToken, checkFieldCreatePost, createPostController)
	router.Get("/:userId", CheckUserIdParams, getPostsByUserIdController)
}
