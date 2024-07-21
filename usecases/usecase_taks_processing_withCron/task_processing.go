package usecase

import (
	"context"
	"log"
	"sync"

	tasks_models "github.com/bricevrs/go-task-light/task"
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
func (proc TaskProcessorPrototype) ExecuteBatch(ctx context.Context, tasks []tasks_models.TaskQuerier) error {
	if len(tasks) == 0 {
		return nil
	}

	var err error

	rounds := len(tasks) / proc.MaxGoroutines
	rest := len(tasks) % proc.MaxGoroutines

	for i := 0; i < rounds; i++ {
		var wg sync.WaitGroup
		wg.Add(proc.MaxGoroutines)
		tasksToExecute := tasks[i*proc.MaxGoroutines : (i+1)*proc.MaxGoroutines]
		for _, task := range tasksToExecute {
			go func(task tasks_models.TaskQuerier) {
				defer wg.Done()
				err = task.Execute(ctx, task)
				if err != nil {
					log.Printf("Error executing task %s: %s", task.GetId(), err.Error())
				}
			}(task)
		}
		wg.Wait()
	}

	if rest > 0 {
		var wg sync.WaitGroup
		wg.Add(rest)
		tasksToExecute := tasks[rounds*proc.MaxGoroutines:]
		for _, task := range tasksToExecute {
			go func(task tasks_models.TaskQuerier) {
				defer wg.Done()
				err = task.Execute(ctx, task)
				if err != nil {
					log.Printf("Error executing task %s: %s", task.GetId(), err.Error())
				}
			}(task)
		}
		wg.Wait()
	}

	return nil
}

func (proc TaskProcessorPrototype) Execute(ctx context.Context, t tasks_models.TaskQuerier) error {
	return t.Execute(ctx, t)
}
