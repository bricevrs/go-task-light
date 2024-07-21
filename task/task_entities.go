package tasks_models

import (
	"time"

	"github.com/google/uuid"
)

/*

The goal of this file is to define what every task should have in common.
For instance, a task should have a status that will indicate if the task is done or not.
A task should also have a payload that will contain the information needed to execute the task.

*/

type TaskRelaunchPolicy string

const (
	TaskRelaunchPolicyNone TaskRelaunchPolicy = "NONE"
)

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusDone       TaskStatus = "DONE"
	TaskStatusError      TaskStatus = "ERROR"
)

type TaskLog struct {
	Datetime time.Time  `json:"datetime"`
	Status   TaskStatus `json:"status"`
	Error    string     `json:"error"`
}

type Task struct {
	ID        uuid.UUID            `json:"id"`
	Payload   TaskPayloadMarshaler `json:"payload"`
	Status    TaskStatus           `json:"status"`
	StatusLog []TaskLog            `json:"status_log"`
	//RelaunchPolicy TaskRelaunchPolicy
}
