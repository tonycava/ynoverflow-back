package comment

import (
	"backend/dto"
	"backend/post"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func checkFieldCreateComment(c *fiber.Ctx) error {
	var comment dto.Comment
	var err = c.BodyParser(&comment)
	var errorValidation = comment.Validate()

	if err != nil || errorValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: utils.Parse(errorValidation.Error()),
			Code:    fiber.StatusBadRequest,
			Data:    nil,
		})
	}

	c.Locals("comment", comment)
	return c.Next()
}

func CheckPostIdParams(c *fiber.Ctx) error {
	var postId = c.Params("postId")
	var postIdFromDb = post.GetPostById(postId)

	if postIdFromDb.ID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.YnoverflowResponse{
			Message: "Post does not exist",
			Code:    fiber.StatusBadRequest,
		})
	}

	return c.Next()
}
