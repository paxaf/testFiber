package repository

import "github.com/jackc/pgx/v5"

type TaskRepository struct {
	Conn *pgx.Conn
}

func (r *TaskRepository) Add(title, desc, status string) {

}
