package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/taskmanage2/internal/config"
	"github.com/priyanshu334/taskmanage2/internal/database"
	"github.com/priyanshu334/taskmanage2/internal/logger"
	"github.com/priyanshu334/taskmanage2/internal/middleware"
	"github.com/priyanshu334/taskmanage2/internal/models/task"
	"github.com/priyanshu334/taskmanage2/internal/models/user"
)

func main() {
	cfg := config.LoadConfig()
	logger.Init()
	log := logger.Log

	database.Connect(cfg)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	app.Use(middleware.Logger(log))
	userRepo := user.NewRepository()
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	api := app.Group("/api")
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	protected := api.Group("/", middleware.AuthMiddleware)
	protected.Get("/me", func(c fiber.Ctx) error {
		userID := c.Locals("user_id")
		return c.JSON(fiber.Map{
			"user_id": userID,
		})
	})
	taskRepo := task.NewRepository()
	taskService := task.NewService(taskRepo)
	taskHandler := task.NewHandler(taskService)
	protected.Post("/tasks", taskHandler.Create)
	protected.Get("/tasks", taskHandler.GetAll)
	protected.Put("/task/:id", taskHandler.Update)
	protected.Delete("/task/:id", taskHandler.Delete)
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Task mange api is running",
		})
	})
	app.Listen(":" + cfg.Port)
}
