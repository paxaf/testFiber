package main

import (
	"context"
	"fmt"
	"os"

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
	conn := repository.ConnectDB()
	defer conn.Close(context.Background())
	// создаём структуру для хранения подключения
	db := repository.TaskRepository{Conn: conn}

	app := fiber.New()
	app.Post("/tasks", handlers.AddTask(db))
	app.Get("/tasks", handlers.GetTask(db))
	app.Delete("/tasks/:id", handlers.DeleteTask(db))
	app.Put("/tasks/:id", handlers.UpdateTask(db))

	app.Listen(fmt.Sprintf(":%s", os.Getenv("TODO_PORT")))
}
