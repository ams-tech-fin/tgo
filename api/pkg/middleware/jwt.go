package middleware

import (
	"tgo/api/internal/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	jwt "github.com/golang-jwt/jwt/v5"

	"time"
)

var jwtSecret = []byte(config.GetEnv("JWT_SECRET", "SUPERSECRETVALUE"))

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtSecret,
		ContextKey:  "user",
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized, invalid or missing token",
			})
		},
	})

}

// Função para gerar o token JWT
func GenerateJWT(userID, userEmail string) (string, error) {
	// Criação das claims (dados do usuário no payload)
	claims := jwt.MapClaims{
		"id":    userID,
		"email": userEmail,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token válido por 24 horas
	}

	// Criando o token com o método de assinatura e as claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assinando o token com a chave secreta
	return token.SignedString(jwtSecret)
}
