package routes

import (
	"golang_redis_pgsql/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/getUsers", func(c *fiber.Ctx) error {
		return handlers.GetAllUsers(c)
	})
	app.Post("/createUsers", func(c *fiber.Ctx) error {
		return handlers.CreateUser(c)
	})
}
