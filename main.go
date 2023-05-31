package main

import (
	"backend/database"
	"backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	database.Connect()

	app.Use(cors.New())

	router.SetupRoutes(app)
}
