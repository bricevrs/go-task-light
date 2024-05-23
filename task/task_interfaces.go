package task

/*

The goal of this file is to define how every task should behave.
For instance, a task should be able to be added to the database, updated, and retrieved.

*/

import (
	"context"
	"encoding/json"
)

type TaskQuerier interface {
	Add(ctx context.Context, arg ...interface{}) error         // Add to the database
	UpdateStatus(status TaskStatus, arg ...interface{}) error  // Update the status of the task in the database
	Execute(context context.Context, arg ...interface{}) error // Execute the task
	//Relaunch() error
}

type TaskPayloadMarshaler interface {
	ToJSON() ([]byte, error)
	FromJSON(json.RawMessage) error
}
