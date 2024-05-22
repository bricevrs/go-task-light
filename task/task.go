package task

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "TODO"
	TaskStatusDo    TaskStatus = "DO"
	TaskStatusDone  TaskStatus = "DONE"
	TaskStatusError TaskStatus = "ERROR"
)

type TaskLog struct {
	UpdatedAt string
	Status    TaskStatus
	Error     error
}

type TaskQuerier interface {
	AddTask(ctx context.Context, arg ...interface{}) error
	UpdateTaskStatus(ctx context.Context, arg ...interface{}) error
	GetTask(ctx context.Context, arg ...interface{}) (*Task, error)
}

type TaskProcessor interface {
	ExecTask(ctx context.Context, arg ...interface{}) error
}

type TaskManager struct {
	db   TaskQuerier
	proc TaskProcessor
}

type Task struct {
	ID        uuid.UUID
	Payload   json.RawMessage
	Status    TaskStatus
	StatusLog []TaskLog
}

func New(db TaskQuerier, proc TaskProcessor) *TaskManager {
	return &TaskManager{
		db:   db,
		proc: proc,
	}
}
