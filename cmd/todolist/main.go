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
	app.Post("/tasks", handlers.AddTask(db))
	app.Get("/tasks", handlers.GetTask(db))
	app.Delete("/tasks/:id", handlers.DeleteTask(db))
	app.Put("/tasks/:id", handlers.UpdateTask(db))
	app.Listen(":8080")
}
