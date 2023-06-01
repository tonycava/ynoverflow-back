package comment

import (
	"backend/post"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesComments(router fiber.Router) {
	router.Get("/post/:postId", CheckPostIdParams, GetCommentsByPostIdController)
	router.Get("/user/:userId", post.CheckUserIdParams, GetCommentsByUserIdController)
	router.Post("/", utils.CheckToken, utils.DecodeToken, CheckFieldCreateComment, CheckIfPostExist, CreateCommentController)
}
