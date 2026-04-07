package task

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/taskmanage2/internal/pkg/response"
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

func (h *Handler) Update(c fiber.Ctx) error {
	taskID := c.Params("id")
	userID := c.Locals("user_id").(string)
	var req UpdateTaskRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, 400, err.Error())
	}
	err := h.service.Update(taskID, userID, &req)
	if err != nil {
		return response.Error(c, 403, err.Error())
	}
	return response.Success(c, "task updated")
}

func (h *Handler) Delete(c fiber.Ctx) error {
	taskID := c.Params("id")
	userID := c.Locals("user_id").(string)
	err := h.service.Delete(taskID, userID)
	if err != nil {
		return response.Error(c, 403, err.Error())
	}
	return response.Success(c, "task deleted")
}
