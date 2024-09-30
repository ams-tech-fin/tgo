package routes

import (
	"tgo/api/internal/di"
	"tgo/api/internal/modules/http/routes/subroutes"

	"github.com/gofiber/fiber/v2"
)

func StartRoutes(app *fiber.App, container *di.Container) {
	subroutes.UserRoutes(app, container)
}
