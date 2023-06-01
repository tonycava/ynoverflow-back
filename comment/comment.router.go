package comment

import (
	"backend/post"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesComments(router fiber.Router) {
	router.Get("/post/:postId", CheckPostIdParams, getCommentsByPostId)
	router.Get("/user/", utils.CheckToken, utils.DecodeToken, getCommentsByUserId)
	router.Post("/", utils.CheckToken, utils.DecodeToken, checkFieldCreateComment, post.CheckIfPostExist, createComment)
}
