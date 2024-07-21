package tasks_processor

import (
	"context"

	tasks_models "github.com/bricevrs/go-task-light/task"
)

type TaskProcessorQuerier interface {
	Execute(ctx context.Context, task tasks_models.TaskQuerier) error
	ExecuteBatch(ctx context.Context, tasks []tasks_models.TaskQuerier) error
	RetrieveTasks(key string) ([]tasks_models.TaskQuerier, error)
}
