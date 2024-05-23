package processor

import (
	"context"

	"github.com/bricevrs/go-task-light/task"
)

type TaskProcessorQuerier interface {
	FetchTasks(ctx context.Context) ([]task.TaskQuerier, error)
	ExecuteTask(ctx context.Context, task task.TaskQuerier) error
	ExecuteBatchTasks(ctx context.Context, tasks []task.TaskQuerier) error
	RelaunchTask(ctx context.Context, task task.TaskQuerier) error
}
