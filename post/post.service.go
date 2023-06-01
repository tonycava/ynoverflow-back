package post

import (
	"backend/database"
	"backend/dto"
	"backend/models"
)

func GetPosts(start, end int) []dto.PostDTO {
	var posts []dto.PostDTO
	database.DB.Db.
		Select("p.id, username, p.title, p.content, p.created_at").
		Table("users u").
		Joins("join posts p on u.id = p.user_id").
		Order("p.created_at ASC").
		Offset(start).
		Limit(end).
		Scan(&posts)
	return posts
}

func GetPostsByUserId(userId string) []dto.PostDTO {
	var posts []dto.PostDTO
	database.DB.Db.
		Select("p.id, username, p.title, p.content, p.created_at").
		Table("users u").
		Joins("join posts p on u.id = p.user_id").
		Order("p.created_at ASC").
		Where("user_id", userId).
		Scan(&posts)
	return posts
}

func CreatePost(post models.Post) models.Post {
	var createdPost models.Post
	database.DB.Db.Create(&post).Scan(&createdPost)
	return createdPost
}

func GetPostById(postId string) models.Post {
	var post models.Post
	database.DB.Db.Table("posts").Select("posts.*, users.username, users.email").Joins("left join users on users.id = posts.user_id").Where("posts.id = ?", postId).Scan(&post)
	return post
}
