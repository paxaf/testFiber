package repository

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://tasks:tasks@localhost:5432/dbname?sslmode=disable")
	if err != nil {
		log.Fatalf("ошибка подключения к БД: %v", err)
	}
	// проверка соединения
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("невозможно установить соединение с БД: %v", err)
	}
	log.Println("Подключение с базой данных установлено")

	filePath := "migration/001_init.sql"
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("ошибка получения рабочего каталога: %v", err)
	}
	sqlFile := path.Join(workDir, filePath)
	sqlBytesFile, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("ошибка чтения файла %s: %v", sqlFile, err)
	}

	sqlReadFile := string(sqlBytesFile)
	_, err = conn.Exec(context.Background(), sqlReadFile)
	if err != nil {
		log.Fatalf("ошибка выполнения SQL запроса: %v", err)
	}

	return conn
}
