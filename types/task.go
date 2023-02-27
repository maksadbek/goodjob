package types

import "time"

type TaskState struct {
	Title       string
	Descriptoin string
}

type Task struct {
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	State       TaskState
}
