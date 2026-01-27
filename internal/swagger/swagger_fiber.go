package swagger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func fiberSwagger() fiber.Handler {
	return swagger.New(swagger.Config{
		Title: "Fiber Production API",
	})
}
