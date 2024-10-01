package subroutes

import (
	"tgo/api/internal/di"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, container *di.Container) {

	authRoutes := app.Group("/auth")
	authRoutes.Post("/login", container.UserController.Auth)
}
