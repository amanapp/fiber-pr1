package user

import (
	"github.com/gofiber/fiber/v2"
	"fiber-app/internal/querybuilder"
)

// ListUsers godoc
// @Summary List users
// @Description Get all users with search
// @Tags Users
// @Accept json
// @Produce json
// @Param search query string false "Search by name"
// @Security BearerAuth
// @Success 200 {array} User
// @Failure 401 {object} map[string]string
// @Router /users [get]
func List(c *fiber.Ctx) error {
	filter := querybuilder.New().
		Regex("name", c.Query("search")).
		Build()

	users, err := ListUsers(filter)
	if err != nil {
		return err
	}

	return c.JSON(users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param body body user.CreateUserDTO true "Create user"
// @Security basicAuth
// @Success 201 {object} User
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /users [post]
func Create(c *fiber.Ctx) error {
	var user CreateUserDTO
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if err := CreateUser(&user); err != nil {
		return err
	}

	return c.JSON(user)
}


