package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/taskmanage2/internal/config"
	"github.com/priyanshu334/taskmanage2/internal/database"
	"github.com/priyanshu334/taskmanage2/internal/logger"
)

func main() {
	cfg := config.LoadConfig()
	logger.Init()
	database.Connect(cfg)

	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Task mange api is running",
		})
	})
	app.Listen(":" + cfg.Port)
}
