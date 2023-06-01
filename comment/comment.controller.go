package comment

import (
	"backend/dto"
	"backend/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCommentController(c *fiber.Ctx) error {
	var comment = c.Locals("comment").(dto.Comment)
	var userId = c.Locals("decodedToken").(*dto.Claims).ID

	createdComment := CreateComment(models.Comment{
		Content: comment.Content,
		UserId:  userId,
		PostId:  comment.PostId,
	})

	return c.Status(fiber.StatusCreated).JSON(createdComment)
}

func GetCommentsByPostIdController(c *fiber.Ctx) error {
	var postId = c.Params("postId")

	comments := GetCommentsByPostId(postId)

	return c.Status(fiber.StatusOK).JSON(comments)
}

func GetCommentsByUserIdController(c *fiber.Ctx) error {
	var userId = c.Params("userId")

	comments := GetCommentsByUserId(userId)

	return c.Status(fiber.StatusOK).JSON(comments)
}
