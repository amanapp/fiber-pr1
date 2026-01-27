package app

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"fiber-app/internal/config"
	"fiber-app/internal/database"
	"fiber-app/internal/error"
	"fiber-app/internal/swagger"
)

func New() *fiber.App {
	config.LoadEnv()
    database.ConnectMongo()

	app := fiber.New(fiber.Config{
		ErrorHandler: errors.Handler,
	})

	registerMiddlewares(app)
	registerRoutes(app)
	swagger.Register(app)
     
	log.Println("Swagger running on http://localhost:3010/swagger/index.html")
	return app
}


