package comment

import (
	"backend/database"
	"backend/models"
)

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

func CreateComment(comment models.Comment) models.Comment {
	var createdComment models.Comment
	database.DB.Db.Create(&comment).Joins("inner join users on users.id = comments.user_id").Scan(&createdComment)
	return createdComment
}
