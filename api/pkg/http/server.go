package http

import (
	"fmt"
	"log"
	"os"
	"tgo/api/internal/di"
	"tgo/api/internal/modules/http/routes"
	"tgo/api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {

	dat, err := os.ReadFile("VERSION")
	if err != nil {
		log.Fatalf("File Version does not exist: %v", err)
	}
	app := fiber.New(fiber.Config{
		Prefork: true,
		AppName: fmt.Sprintf("TGo v%v", string(dat)),
	})

	app.Use(middleware.CORS())
	app.Use(middleware.Helmet())

	container, err := di.NewContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	routes.StartRoutes(app, container)

	return app
}
