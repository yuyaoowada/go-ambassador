package main

import (
	"ambassador/src/database"
	"ambassador/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
