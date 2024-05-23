package main

import (
	"context"
	"log"

	"github.com/bricevrs/go-task-light/task"
	usecase "github.com/bricevrs/go-task-light/usecases/usecase_taks_processing_withCron"
)

func retrieveTasks() []task.TaskQuerier {
	var tasks []task.TaskQuerier = make([]task.TaskQuerier, 0)
	for i := 0; i < 10; i++ {
		// task can be of a generic type
		// we can convert it to a specific type
		t, err := usecase.NewWaterPlantTask("Plant1", 10)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, t)
	}
	return tasks
}

func main() {
	ctx := context.Background()
	taskProcessor := usecase.NewTaskProcessorPrototype(10)

	// 1 - Retrieve the tasks with the status "todo"
	var tasks []task.TaskQuerier = retrieveTasks()

	// 3 - Execute the tasks with the task processor
	err := taskProcessor.ExecuteBatch(ctx, tasks)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("All Tasks executed")
	}

	// 4 - Execute a specific task
	err = taskProcessor.Execute(ctx, tasks[0])
	log.Println(err)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Task 0 executed")
	}
}
