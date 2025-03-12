package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/paxaf/testFiber/internal/models"
)

type TaskRepository struct {
	Conn *pgx.Conn
}

func (r *TaskRepository) Add(t models.Task) error {
	switch {
	case t.Status == "in_progress" || t.Status == "done":
		_, err := r.Conn.Exec(context.Background(), "INSERT INTO Tasks (title, description, status) VALUES ($1, $2, $3);",
			t.Title,
			t.Description,
			t.Status)
		if err != nil {
			log.Println(err)
			return err
		}
	default:
		_, err := r.Conn.Exec(context.Background(), "INSERT INTO Tasks (title, description) VALUES ($1, $2);",
			t.Title,
			t.Description)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (r *TaskRepository) Get() ([]models.Task, error) {
	var tasks []models.Task
	rows, err := r.Conn.Query(context.Background(), `SELECT id, title, description, status, created_at::TEXT AS created_at_text, updated_at ::TEXT AS updated_at_text
	FROM Tasks`)
	if err != nil {
		return nil, fmt.Errorf("ошибка с базой данных: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var task models.Task
		var id int
		err := rows.Scan(
			&id,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("ошибка инициализации строки: %w", err)
		}
		task.Id = strconv.Itoa(id)
		if task.CreatedAt == task.UpdatedAt {
			task.UpdatedAt = ""
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации строк: %w", err)
	}
	return tasks, nil
}
