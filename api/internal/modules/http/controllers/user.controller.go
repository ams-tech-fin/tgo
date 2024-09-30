package controller

import (
	service "tgo/api/internal/services/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	type Request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	user, err := c.service.CreateUser(ctx.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	type userReponse struct {
		ID    string
		Name  string
		Email string
	}
	var dest userReponse

	errorcopy := copier.Copy(&dest, &user)
	if errorcopy != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errorcopy.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(dest)
}

func (c *UserController) GetUserByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	user, err := c.service.GetUserByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	type userReponse struct {
		ID    string
		Name  string
		Email string
	}
	var dest userReponse

	errorcopy := copier.Copy(&dest, &user)
	if errorcopy != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errorcopy.Error()})
	}

	return ctx.JSON(dest)
}

func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {

	user, err := c.service.GetAllUsers(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	type userReponse struct {
		ID    string
		Name  string
		Email string
	}
	var dest []userReponse

	errorcopy := copier.Copy(&dest, &user)
	if errorcopy != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errorcopy.Error()})
	}

	return ctx.JSON(dest)
}
