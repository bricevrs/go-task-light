package task

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
	Payload        TaskPayloadMarshaler
	Status         TaskStatus
	StatusLog      []TaskLog
	RelaunchPolicy TaskRelaunchPolicy
}
