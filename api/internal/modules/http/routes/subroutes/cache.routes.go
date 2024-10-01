package subroutes

import (
	"tgo/api/internal/di"

	"github.com/gofiber/fiber/v2"
)

func CacheRoutes(app *fiber.App, container *di.Container) {

	userRoutes := app.Group("/cache")
	userRoutes.Post("/", container.CacheController.CacheTest)
	userRoutes.Delete("/:id", container.CacheController.CacheTest)
	userRoutes.Get("/:id", container.CacheController.CacheTest)
}
