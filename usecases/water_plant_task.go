package usecases

import (
	"encoding/json"

	"github.com/bricevrs/go-task-light/task"
)

type WaterPlantTask struct {
	task.Task
}

type WaterPlantPayload struct {
	PlantName     string `json:"plantName"`
	WaterQuantity int    `json:"waterQuantity"`
}

func NewWaterPlantTask(pn string, wq int) (*WaterPlantTask, error) {
	payload := WaterPlantPayload{
		PlantName:     pn,
		WaterQuantity: wq,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return &WaterPlantTask{
		Task: task.Task{
			Payload:  payloadBytes,
			Status:   task.TaskStatusTodo,
			Relaunch: task.TaskRelaunchPolicyNone,
		},
	}, nil
}

func (t *WaterPlantTask) AddTask() error {
	// to be implemented
	return nil
}
func (t *WaterPlantTask) UpdateTaskStatus() error {
	// to be implemented
	return nil
}
func (t *WaterPlantTask) GetTask() (*WaterPlantTask, error) {
	// to be implemented
	return nil, nil
}
func (t *WaterPlantTask) MarshalTaskPayload(wpp WaterPlantPayload) (json.RawMessage, error) {
	return json.Marshal(wpp)
}
func (t *WaterPlantTask) UnmarshalTaskPayload(raw json.RawMessage) (WaterPlantPayload, error) {
	var wpp WaterPlantPayload
	err := json.Unmarshal(raw, &wpp)
	return wpp, err
}
