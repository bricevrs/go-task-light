package usecase

import (
	"context"
	"fmt"

	model "github.com/bricevrs/go-task-light/task"
	processor "github.com/bricevrs/go-task-light/task_processor"
)

type TaskProcessorPrototype struct {
	processor.TaskProcessor
}

func NewTaskProcessorPrototype(maxGoroutines int) *TaskProcessorPrototype {
	return &TaskProcessorPrototype{
		TaskProcessor: processor.TaskProcessor{
			MaxGoroutines: maxGoroutines,
		},
	}
}

func (proc TaskProcessorPrototype) ExecuteBatch(ctx context.Context, tasks []model.TaskQuerier) error {
	// Can add some logic here to handle the execution of the tasks
	// For instance, we can add a retry mechanism
	// Or we can add a timeout

	// We use goroutines to execute the tasks in parallel
	// We use a worker pool to limit the number of tasks executed in parallel

	fmt.Println("Executing tasks...")
	fmt.Println("Number of tasks to execute: ", len(tasks))
	fmt.Println("Number of goroutines: ", proc.MaxGoroutines)

	return nil
}

func (proc TaskProcessorPrototype) Execute(ctx context.Context, t model.TaskQuerier) error {
	// Can add some logic here to handle the execution of the task
	// For instance, we can add a retry mechanism
	// Or we can add a timeout

	err := t.Execute(ctx)
	if err != nil {
		return err
	}

	return nil
}
