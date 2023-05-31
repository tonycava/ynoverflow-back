package database

import (
	"backend/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB *DBInstance

func Connect() {

	pgUser, _ := os.LookupEnv("POSTGRES_USER")
	pgPassword, _ := os.LookupEnv("POSTGRES_PASSWORD")
	pgDbname, _ := os.LookupEnv("POSTGRES_DB")

	fmt.Println(pgUser, pgPassword, pgDbname)

	db, err := gorm.Open(postgres.Open("host=localhost user="+pgUser+" password="+pgPassword+" dbname="+pgDbname+" port=10001 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	DB = &DBInstance{Db: db}
}
