package database

import (
	"backend/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func Connect() {
	pgDatabase, _ := os.LookupEnv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(pgDatabase), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	fmt.Println("Database migrated")

	DB = DBInstance{
		Db: db,
	}
}
