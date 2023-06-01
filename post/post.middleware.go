package post

import (
	"backend/dto"
	"backend/user"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func checkFieldCreatePost(c *fiber.Ctx) error {
	var post dto.Post
	err := c.BodyParser(&post)
	var errorValidation = post.Validate()

	if err != nil || errorValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: utils.Parse(errorValidation.Error()),
			Code:    fiber.StatusBadRequest,
			Data:    nil,
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

func CheckIfPostExist(c *fiber.Ctx) error {
	var comment = c.Locals("comment").(dto.Comment)
	var postId = comment.PostId
	var postIdFromDb = GetPostById(postId)

	if postIdFromDb.ID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: "Post does not exist",
			Code:    fiber.StatusBadRequest,
		})
	}

	return c.Next()
}
