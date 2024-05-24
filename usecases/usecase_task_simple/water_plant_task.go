package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	model "github.com/bricevrs/go-task-light/task"
)

type Dao interface {
	waterThePlant(plantName string, quantityOfWater int) error
}

type DaoImpl struct {
}

func (dao DaoImpl) waterThePlant(plantName string, quantityOfWater int) error {
	log.Println("Watering the plant...")
	return nil
}

type DB interface {
	save() error
	setStatus(status model.TaskStatus) error
}

type DBImpl struct {
}

func (db DBImpl) save() error {
	fmt.Println("Saving to the database...")
	return nil
}
func (db DBImpl) setStatus(status model.TaskStatus) error {
	fmt.Println("Setting the status in the database...")
	return nil
}

type WaterPlantTask struct {
	model.Task
}

type WaterPlantPayload struct {
	PlantName     string `json:"plantName"`
	WaterQuantity int    `json:"waterQuantity"`
}

func NewWaterPlantTask(plantName string, waterQuantity int) (*WaterPlantTask, error) {
	t := &WaterPlantTask{
		Task: model.Task{
			Status:  model.TaskStatusTodo,
			Payload: nil,
		},
	}

	payload := WaterPlantPayload{
		PlantName:     plantName,
		WaterQuantity: waterQuantity,
	}

	t.Payload = &payload

	return t, nil
}

func (t WaterPlantTask) Add(ctx context.Context, arg ...interface{}) error {
	db := arg[0].(DB)
	err := db.save()
	if err != nil {
		return err
	}
	return nil
}
func (t WaterPlantTask) UpdateStatus(status model.TaskStatus, arg ...interface{}) error {
	db := arg[0].(DB)
	err := db.setStatus(status)
	if err != nil {
		return err
	}
	return nil
}

func (t WaterPlantTask) Execute(ctx context.Context, arg ...interface{}) error {
	dao := arg[0].(Dao)
	payload := t.Payload
	if payload == nil {
		return fmt.Errorf("payload is nil")
	}
	log.Println("Executing the task...")
	log.Println("Plant Name:", payload.(*WaterPlantPayload).PlantName)
	log.Println("Water Quantity:", payload.(*WaterPlantPayload).WaterQuantity)

	// send the payload to the dao
	err := dao.waterThePlant(payload.(*WaterPlantPayload).PlantName, payload.(*WaterPlantPayload).WaterQuantity)
	if err != nil {
		return err
	}

	// update the status
	t.UpdateStatus(model.TaskStatusDone)
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
