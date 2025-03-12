package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

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
	rows, err := r.Conn.Query(context.Background(), `SELECT id, title, description, status, created_at, updated_at
	FROM Tasks`)
	if err != nil {
		return nil, fmt.Errorf("ошибка с базой данных: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var task models.Task
		var createdAt time.Time
		var updatedAt time.Time
		var id int
		err := rows.Scan(
			&id,
			&task.Title,
			&task.Description,
			&task.Status,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("ошибка инициализации строки: %w", err)
		}
		task.Id = strconv.Itoa(id)
		task.CreatedAt = createdAt.Format(models.OutputFormat)
		task.UpdatedAt = updatedAt.Format(models.OutputFormat)
		if createdAt == updatedAt {
			task.UpdatedAt = ""
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации строк: %w", err)
	}
	return tasks, nil
}

func (r *TaskRepository) Update(t models.Task, id int) error {

	// у меня тут на '!=' ругался линтер
	if !(t.Status == "done" || t.Status == "in_progress") {
		row := r.Conn.QueryRow(context.Background(), "SELECT status FROM Tasks WHERE id = $1", id)
		err := row.Scan(&t.Status)
		if err != nil {
			return err
		}
	}

	affectedRows, err := r.Conn.Exec(context.Background(), "UPDATE Tasks SET title = $1, description = $2, status = $3, updated_at = now() WHERE id = $4",
		t.Title,
		t.Description,
		t.Status,
		id)

	if err != nil {
		return err
	}
	if affectedRows.RowsAffected() < 1 {
		return fmt.Errorf("ошибка: ни одна строка не была изменена")
	}
	return nil
}

func (r *TaskRepository) Delete() {}
