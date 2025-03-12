package models

import "strings"

type Task struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type TasksResponse struct {
	Tasks []Task `json:"tasks"`
}

func (t *Task) IsValid() bool {
	return len(strings.TrimSpace(t.Title)) > 0
}

const InputFormat string = "2006-01-02 15:04:05.999999"
const OutputFormat string = "02/01/2006 15:04"
