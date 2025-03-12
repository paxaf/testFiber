package handlers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/paxaf/testFiber/internal/models"
	"github.com/paxaf/testFiber/internal/repository"
)

func AddTask(repo repository.TaskRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		params := new(models.Task)
		if err := c.QueryParser(params); err != nil {
			log.Printf("ошибка парсинга параметров: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "ошибка запроса",
			})
		}
		if !params.IsValid() {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "поле Title не может быть пустым",
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func UpdateTask(db repository.TaskRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString(fmt.Sprintf("Task %s updated", id))
	}
}

func DeleteTask(db repository.TaskRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString(fmt.Sprintf("Task %s deleted", id))
	}
}

func GetTask(db repository.TaskRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Task got")
	}
}
