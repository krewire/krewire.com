package models

import (
	"time"
)

// Task is one row of the tasks table.
type Task struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Done      bool      `json:"done" db:"done"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// TableName maps Task to the "tasks" table (pkg/orm would default to "task").
func (Task) TableName() string { return "tasks" }
