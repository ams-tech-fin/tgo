package controller

import (
	"fmt"
	service "tgo/api/internal/services/user"
	"tgo/api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

func (c *UserController) Auth(ctx *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	user, err := c.service.AuthUser(ctx.Context(), req.Email, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	type userReponse struct {
		AccessToken string `json:"access_token"`
	}
	var dest userReponse

	var userId string = user.ID.String()
	dest.AccessToken, _ = middleware.GenerateJWT(userId, user.Email)

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

	// TODO: Maybe one middleware to use as global value?
	token := ctx.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["id"])

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
