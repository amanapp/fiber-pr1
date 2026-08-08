package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func registerMiddlewares(app *fiber.App) {

	// PANIC RECOVERY (CRITICAL)
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true, // dev me useful
	}))

	// REQUEST LOGGER
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} - ${status} - ${method} ${path} ${latency}\n",
	}))

}

// thsi go middle ware
