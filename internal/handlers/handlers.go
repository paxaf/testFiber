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

		err := repo.Add(*params)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "ошибка добавления задачи",
			})
		}
		return c.SendStatus(fiber.StatusOK)
	}
}

func UpdateTask(repo repository.TaskRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString(fmt.Sprintf("Task %s updated", id))
	}
}

func DeleteTask(repo repository.TaskRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString(fmt.Sprintf("Task %s deleted", id))
	}
}

func GetTask(repo repository.TaskRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tasks, err := repo.Get()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("ошибка: %v", err),
			})
		}
		if len(tasks) == 0 {
			return c.JSON(map[string]string{})
		} else {
			return c.JSON(tasks)
		}
	}
}
