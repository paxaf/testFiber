package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/paxaf/testFiber/internal/handlers"
	"github.com/paxaf/testFiber/internal/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		//	log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}
	db := repository.ConnectDB()
	defer db.Close(context.Background())

	app := fiber.New()
	app.Post("/tasks", handlers.AddTask)
	app.Get("/tasks", handlers.GetTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Listen(":8080")
}
