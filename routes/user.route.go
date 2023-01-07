package routes

import (
	"github.com/amm834/fiber-modular-rest-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	//app.Get("/users", controllers.GetUsers)
	app.Get("/users/:id", controllers.GetUser)
	app.Post("/users", controllers.CreateUser)
	//app.Delete("/users/:id", controllers.DeleteUser)
}
