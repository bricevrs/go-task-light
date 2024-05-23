package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bricevrs/go-task-light/task"
)

type WaterPlantTask struct {
	task.Task
	// can add a db sql connection here
}

type WaterPlantPayload struct {
	PlantName     string `json:"plantName"`
	WaterQuantity int    `json:"waterQuantity"`
}

func NewWaterPlantTask(plantName string, waterQuantity int) (*WaterPlantTask, error) {
	t := &WaterPlantTask{
		Task: task.Task{
			Status:         task.TaskStatusTodo,
			RelaunchPolicy: task.TaskRelaunchPolicyNone,
			Payload:        nil,
		},
	}

	payload := WaterPlantPayload{
		PlantName:     plantName,
		WaterQuantity: waterQuantity,
	}

	t.Payload = &payload

	return t, nil
}

func (t WaterPlantTask) Add() error {
	// to be implemented
	return nil
}
func (t WaterPlantTask) UpdateStatus(status task.TaskStatus) error {
	// to be implemented
	return nil
}

func (t WaterPlantTask) Execute() error {
	// to be implemented
	return nil
}

func (t WaterPlantTask) Relaunch() error {
	// to be implemented
	return nil
}

func (p WaterPlantPayload) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *WaterPlantPayload) FromJSON(raw json.RawMessage) error {
	err := json.Unmarshal(raw, &p)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	plantName := "Aloe Vera"
	waterQuantity := 100
	fmt.Println("...Creating a new WaterPlantTask...")
	fmt.Println("Plant Name:", plantName, "\nWater Quantity:", waterQuantity)
	task, err := NewWaterPlantTask(plantName, waterQuantity)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task created successfully\n")
	fmt.Println("Task Status:", task.Status)
	fmt.Println("Task Relaunch Policy:", task.RelaunchPolicy)
	fmt.Println("Task Payload:", task.Payload)

	fmt.Println("\n...Marshalling the payload...")
	payload, err := task.Payload.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Payload marshalled successfully\n")
	fmt.Println("Payload marshalled to string:", string(payload))

	fmt.Println("...Unmarshalling the payload...")
	var payloadUnmarshalled WaterPlantPayload
	err = payloadUnmarshalled.FromJSON(payload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Payload unmarshalled :", payloadUnmarshalled)
}
