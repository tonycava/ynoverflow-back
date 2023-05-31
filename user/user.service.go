package user

import (
	"backend/database"
	"backend/models"
)

func GetUserByEmail(email string) models.User {
	var user models.User
	database.DB.Db.Table("users").Where("email = ?", email).Scan(&user)
	return user
}

func CreateUser(user models.User) models.User {
	var createdUser models.User
	database.DB.Db.Create(&user).Scan(&createdUser)
	return createdUser
}
