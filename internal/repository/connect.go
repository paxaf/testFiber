package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://tasks:tasks@localhost:5432/dbname?sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("Невозможно установить соединение с БД: %v", err)
	}
	log.Println("Подключение с базой данных установлено")
	return conn
}
