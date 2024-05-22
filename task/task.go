package task

import (
	"encoding/json"
)

/*

The goal of this file is to define what every task should have in common.
The properties and the behavior via the TaskQuerier interface.
For example, a task should have a status that will indicate if the task is done or not.
A task should also have a payload that will contain the information needed to execute the task.
A task should be able to be added to the database, updated, and retrieved.

*/

type TaskRelaunchPolicy string

const (
	TaskRelaunchPolicyNone TaskRelaunchPolicy = "NONE"
)

type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "TODO"
	TaskStatusDo    TaskStatus = "DO"
	TaskStatusDone  TaskStatus = "DONE"
	TaskStatusError TaskStatus = "ERROR"
)

type TaskLog struct {
	UpdatedAt string
	Status    TaskStatus
	Error     error
}

type Task struct {
	Payload   json.RawMessage
	Status    TaskStatus
	StatusLog []TaskLog
	Relaunch  TaskRelaunchPolicy
}

type TaskQuerier interface {
	AddTask(arg ...interface{}) error
	UpdateTaskStatus(arg ...interface{}) error
	GetTask(arg ...interface{}) (*Task, error)
	MarshalTaskPayload(arg ...interface{}) (json.RawMessage, error)
	UnmarshalTaskPayload(arg ...interface{}) (interface{}, error)
	ExecuteTask(arg ...interface{}) error
	RelaunchTask(arg ...interface{}) error
}
