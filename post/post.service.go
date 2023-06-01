package post

import (
	"backend/database"
	"backend/models"
)

func GetPosts() []models.Post {
	var posts []models.Post
	database.DB.Db.Find(&posts)
	return posts
}

func GetPostsByUserId(userId string) []models.Post {
	var posts []models.Post
	database.DB.Db.Where("user_id = ?", userId).Find(&posts)
	return posts
}

func CreatePost(post models.Post) models.Post {
	var createdPost models.Post
	database.DB.Db.Create(&post).Scan(&createdPost)
	return createdPost
}
