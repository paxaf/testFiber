package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func AddTask(db *pgx.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Task added")
	}
}

func UpdateTask(db *pgx.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString(fmt.Sprintf("Task %s updated", id))
	}
}

func DeleteTask(db *pgx.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString(fmt.Sprintf("Task %s deleted", id))
	}
}

func GetTask(db *pgx.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Task got")
	}
}
