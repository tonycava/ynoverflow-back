package post

import (
	"backend/dto"
	"backend/user"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func CheckFieldCreatePost(c *fiber.Ctx) error {
	var checkFieldPostArray = []string{"title", "content"}
	var post dto.Post
	err := c.BodyParser(&post)

	if (err != nil) ||
		!utils.CheckFieldPost(post, checkFieldPostArray) {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: "Invalid request body",
			Code:    fiber.StatusBadRequest,
		})
	}

	c.Locals("post", post)
	return c.Next()
}

func CheckUserIdParams(c *fiber.Ctx) error {
	var userId = c.Params("userId")
	var userIdFromDb = user.GetUserById(userId)

	if userIdFromDb.ID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: "User does not exist",
			Code:    fiber.StatusBadRequest,
		})
	}

	return c.Next()
}
