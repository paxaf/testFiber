package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AddTask(c *fiber.Ctx) error {
	return c.SendString("Task added")
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Task %s updated", id))
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Task %s deleted", id))
}

func GetTask(c *fiber.Ctx) error {
	return c.SendString("Task got")
}
