package main

import (
	"context"
	"log"

	model "github.com/bricevrs/go-task-light/task"
	usecase "github.com/bricevrs/go-task-light/usecases/usecase_taks_processing_withCron"
)

func retrieveTasks() []model.TaskQuerier {
	var tasks []model.TaskQuerier = make([]model.TaskQuerier, 0)
	for i := 0; i < 10; i++ {
		// task can be of a generic type
		// we need to convert it to a specific type if we fetch it from the database
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

	// Retrieve the tasks with the status "todo"
	var tasks []model.TaskQuerier = retrieveTasks()

	// Execute the tasks with the task processor
	err := taskProcessor.ExecuteBatch(ctx, tasks)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("All Tasks executed")
	}

	// Execute a specific task
	err = taskProcessor.Execute(ctx, tasks[0])
	log.Println(err)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Task 0 executed")
	}
}
