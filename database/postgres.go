package database

import (
	"backend/user"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), "10001")

	log.Print("Connecting to database...")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database !\n", err)
		os.Exit(2)
	}

	log.Print("Database connected !")

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Print("Migrating database...")
	DB.AutoMigrate(&user.User{}, &user.Claims{})
}
