package user

import (
	"github.com/gofiber/fiber/v2"
	"fiber-app/internal/middleware"
)

func Routes(r fiber.Router) {
	r.Post("/users", middleware.BasicAuth(), Create)
	r.Get("/users", middleware.JWTAuth(), List)
}
