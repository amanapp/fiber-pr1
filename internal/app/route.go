package app

import (
	"github.com/gofiber/fiber/v2"
	"fiber-app/internal/modules/user"
)

func registerRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	user.Routes(api)
}