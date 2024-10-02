package http

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"tgo/api/internal/di"
	"tgo/api/internal/modules/http/routes"
	"tgo/api/internal/modules/queue"
	"tgo/api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(rabbitAdapter *queue.RabbitMQAdapter) *fiber.App {

	dat, err := os.ReadFile("VERSION")
	if err != nil {
		log.Fatalf("File Version does not exist: %v", err)
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
		AppName: fmt.Sprintf("TGo v%v", string(dat)),
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			if fiberErr, ok := err.(*fiber.Error); ok && fiberErr.Code == fiber.StatusNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error":   true,
					"message": "Página não encontrada",
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Erro interno no servidor",
			})
		},
	})

	app.Use(middleware.CORS())
	app.Use(middleware.Helmet())

	container, err := di.NewContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	routes.StartRoutes(app, container)

	relativePath := "./api/assets/favicon.ico"
	absolutePath, _ := filepath.Abs(relativePath)
	app.Static("/favicon.ico", absolutePath)

	app.Post("/publish", func(c *fiber.Ctx) error {
		type Request struct {
			Message string `json:"message"`
		}

		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request",
			})
		}

		// Publicar a mensagem no RabbitMQ

		if err := rabbitAdapter.Publish("test-queue", []byte(req.Message)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to publish message",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Mensagem publicada com sucesso!",
		})
	})

	go func() {
		err := rabbitAdapter.Consume("test-queue", func(message []byte) {
			log.Printf("Mensagem consumida: %s", message)
		})
		if err != nil {
			log.Fatalf("Erro ao consumir mensagens: %v", err)
		}
	}()

	return app
}
