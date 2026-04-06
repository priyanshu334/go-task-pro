package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/taskmanage2/internal/config"
	"github.com/priyanshu334/taskmanage2/internal/pkg/utils"
)

func AuthMiddleware(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "missing token"})
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	cfg := config.LoadConfig()

	claims, err := utils.ValidateToken(tokenStr, cfg.JWTSecret)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
	}

	// Inject user ID into context
	c.Locals("user_id", claims.UserID)

	return c.Next()
}
