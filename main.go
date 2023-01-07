package main

import (
	"github.com/amm834/fiber-modular-rest-api/configs"
	"github.com/amm834/fiber-modular-rest-api/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func init() {
	configs.ConnectDB()
}

func main() {
	app := fiber.New()

	// Routes
	routes.UserRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
