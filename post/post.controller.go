package post

import (
	"backend/dto"
	"backend/models"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func createPostController(c *fiber.Ctx) error {
	var post = c.Locals("post").(dto.Post)
	var userId = c.Locals("decodedToken").(*dto.Claims).ID

	createdPost := CreatePost(models.Post{
		Title:   post.Title,
		Content: post.Content,
		UserId:  userId,
	})

	return c.Status(fiber.StatusCreated).JSON(dto.YnoverflowResponse{
		Message: "Created successfully",
		Code:    fiber.StatusCreated,
		Data:    createdPost,
	})
}

func getPostsByUserIdController(c *fiber.Ctx) error {
	var userId = c.Params("userId")

	return c.Status(fiber.StatusOK).JSON(dto.YnoverflowResponse{
		Message: "Fetched successfully",
		Code:    fiber.StatusOK,
		Data:    GetPostsByUserId(userId),
	})
}

func GetAllPostsController(c *fiber.Ctx) error {
	var start = c.Query("start")
	var end = c.Query("end")

	if start == "" {
		start = "0"
	}
	if end == "" {
		end = "10"
	}

	return c.Status(fiber.StatusOK).JSON(dto.YnoverflowResponse{
		Message: "Fetched successfully",
		Code:    fiber.StatusOK,
		Data:    GetPosts(utils.ParseInt(start), utils.ParseInt(end)),
	})
}
