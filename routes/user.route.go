package routes

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
	app.Get("/users/:id", controllers.GetUser)
	app.Post("/users", controllers.NewUser)
	app.Delete("/users/:id", controllers.DeleteUser)
}
