package main

import (
	"backend/database"
	"backend/router"
	"backend/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	database.Connect()

	app.Use(cors.New())

	router.SetupRoutes(app)

	var PORT = fmt.Sprintf(":%s", utils.GetEnv("PORT"))
	log.Fatal(app.Listen(PORT))
}
