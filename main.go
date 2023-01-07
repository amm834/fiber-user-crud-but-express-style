package main

import (
	"github.com/amm834/fiber-modular-rest-api/configs"
	"github.com/gofiber/fiber/v2"
	"log"
)

func init() {
	configs.ConnectDB()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	log.Fatal(app.Listen(":8000"))
}
