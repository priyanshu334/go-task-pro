package user

import (
	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/taskmanage2/internal/config"
	"github.com/priyanshu334/taskmanage2/internal/pkg/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Register(c fiber.Ctx) error {
	var req RegisterRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.Register(&req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "user Registed"})
}

func (h *Handler) Login(c fiber.Ctx) error {
	var req LoginRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := h.service.Login(&req)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	cfg := config.LoadConfig()
	token, err := utils.GenerateToken(user.ID.String(), cfg.JWTSecret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "token error"})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
