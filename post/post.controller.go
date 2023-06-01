package post

import (
	"backend/dto"
	"backend/models"
	"github.com/gofiber/fiber/v2"
)

func CreatePostController(c *fiber.Ctx) error {
	var post = c.Locals("post").(dto.Post)
	var userId = c.Locals("decodedToken").(*dto.Claims).ID

	createdPost := CreatePost(models.Post{
		Title:   post.Title,
		Content: post.Content,
		UserId:  userId,
	})

	return c.Status(fiber.StatusCreated).JSON(createdPost)
}

func GetPostsByUserIdController(c *fiber.Ctx) error {
	var userId = c.Params("userId")

	posts := GetPostsByUserId(userId)

	return c.Status(fiber.StatusOK).JSON(posts)
}

func GetAllPostsController(c *fiber.Ctx) error {
	posts := GetPosts()

	return c.Status(fiber.StatusOK).JSON(posts)
}
