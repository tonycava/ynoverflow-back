package user

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
	"path/filepath"
)

func GetUserById(userId string) models.User {
	var user models.User
	database.DB.Db.Table("users").Where("id = ?", userId).Scan(&user)
	return user
}

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

func uploadProfilePictureToS3Bucket(file *multipart.FileHeader, userId string) {
	openFile, err := file.Open()
	if err != nil {
		return
	}
	defer openFile.Close()

	// Create a new AWS session with static credentials
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(utils.GetEnv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(utils.GetEnv("AWS_ACCESS_KEY_ID"), utils.GetEnv("AWS_SECRET_ACCESS_KEY"), ""),
	})

	svc := s3.New(session)

	var extension = filepath.Ext(file.Filename)
	params := &s3.PutObjectInput{
		Bucket: aws.String(utils.GetEnv("AWS_BUCKET_NAME")),
		Key:    aws.String("profile-pictures/" + userId + extension),
		Body:   openFile,
	}

	_, err = svc.PutObject(params)
	if err != nil {
		// Handle upload error
		panic(err)
	}

}
