package subroutes

import (
	"tgo/api/internal/di"
	"tgo/api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, container *di.Container) {

	userRoutes := app.Group("/users", middleware.JWTMiddleware())
	userRoutes.Post("/", container.UserController.CreateUser)
	userRoutes.Get("/", container.UserController.GetAllUsers)
	userRoutes.Get("/:id", container.UserController.GetUserByID)
}
