package task

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Create(c fiber.Ctx) error {
	var req CreateTaskRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	userID := c.Locals("user_id").(string)

	if err := h.service.Create(&req, userID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "task created"})
}

func (h *Handler) GetAll(c fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	status := c.Query("status")
	search := c.Query("search")

	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	tasks, err := h.service.GetAll(userID, status, search, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(tasks)
}
