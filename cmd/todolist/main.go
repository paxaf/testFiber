package main

import (
	"context"

	"github.com/paxaf/testFiber/internal/repository"
)

func main() {
	db := repository.ConnectDB()
	defer db.Close(context.Background())
}
