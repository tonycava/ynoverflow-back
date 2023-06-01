package comment

import (
	"backend/dto"
	"backend/models"
	"github.com/gofiber/fiber/v2"
)

func createComment(c *fiber.Ctx) error {
	var comment = c.Locals("comment").(dto.Comment)
	var userId = c.Locals("decodedToken").(*dto.Claims).ID

	createdComment := CreateComment(models.Comment{
		Content: comment.Content,
		UserId:  userId,
		PostId:  comment.PostId,
	})

	return c.Status(fiber.StatusCreated).JSON(dto.YnoverflowResponse{
		Message: "Created successfully",
		Code:    fiber.StatusCreated,
		Data:    createdComment,
	})
}

func getCommentsByPostId(c *fiber.Ctx) error {
	var postId = c.Params("postId")

	comments := GetCommentsByPostId(postId)

	return c.Status(fiber.StatusOK).JSON(dto.YnoverflowResponse{
		Message: "Fetched successfully",
		Code:    fiber.StatusOK,
		Data:    comments,
	})
}

func getCommentsByUserId(c *fiber.Ctx) error {
	var userId = c.Locals("decodedToken").(*dto.Claims).ID

	comments := GetCommentsByUserId(userId)

	return c.Status(fiber.StatusOK).JSON(dto.YnoverflowResponse{
		Message: "Fetched successfully",
		Code:    fiber.StatusOK,
		Data:    comments,
	})
}
