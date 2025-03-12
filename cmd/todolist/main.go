package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/paxaf/testFiber/internal/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		//	log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}
	db := repository.ConnectDB()
	defer db.Close(context.Background())
}
