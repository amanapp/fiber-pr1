package errors

import (
	stdErrors "errors"
	"fiber-app/internal/config"
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if stdErrors.As(err, &e) {
		code = e.Code
	}

	if config.Config.Env == "production" {
		message := "Something went wrong"
		if e != nil {
			message = e.Message
		}
		return c.Status(code).JSON(fiber.Map{
			"message": message,
		})
	}

	response := fiber.Map{
		"error":  err.Error(),
		"path":   c.Path(),
		"method": c.Method(),
		"query":  c.Queries(),
		"body":   string(c.Body()),
	}

	if c.Route() != nil {
		response["params"] = c.AllParams()
	}

	return c.Status(code).JSON(response)
}