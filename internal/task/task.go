package task

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	Deadline  time.Time `json:"deadline,omitempty"`
}

type TaskStore interface {
	Add(task Task) error
	List() ([]Task, error)
	MarkDone(id int) error
	Remove(id int) error
}
