package swagger

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger())
}
