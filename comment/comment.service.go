package comment

import (
	"backend/database"
	"backend/models"
)

type CommentDTO struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Content   string `json:"content"`
	Username  string `json:"username"`
}

func GetCommentsByUserId(userId string) []models.Comment {
	var comments []models.Comment
	database.DB.Db.Table("comments").Select("comments.*, users.username, users.email").Joins("left join users on users.id = comments.user_id").Where("comments.user_id = ?", userId).Scan(&comments)
	return comments
}

func GetCommentsByPostId(postId string) []models.Comment {
	var comments []models.Comment
	database.DB.Db.Table("comments").Select("comments.*, users.username, users.email").Joins("left join users on users.id = comments.user_id").Where("comments.post_id = ?", postId).Scan(&comments)
	return comments
}

func CreateComment(comment models.Comment) CommentDTO {
	var createdComment CommentDTO
	database.DB.Db.
		Create(&comment).
		Joins("inner join users on users.id = comments.user_id").
		Select("comments.*, users.username").
		Scan(&createdComment)
	return createdComment
}
