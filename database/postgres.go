package database

import (
	"backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB *DBInstance

func Connect() {

	pgDatabase, _ := os.LookupEnv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(pgDatabase), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	DB = &DBInstance{Db: db}
}
