package tasks_models

/*

The goal of this file is to define how every task should behave.
For instance, a task should be able to be added to the database, updated, and retrieved.

*/

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type TaskQuerier interface { // Get the ID of the task
	GetId() uuid.UUID
	Add(ctx context.Context, arg ...interface{}) error                                           // Add to the database
	UpdateStatus(ctx context.Context, status TaskStatus, error string, arg ...interface{}) error // Update the status of the task in the database
	Execute(ctx context.Context, arg ...interface{}) error                                       // Execute the task
	//Relaunch() error
}

type TaskPayloadMarshaler interface {
	ToJSON() ([]byte, error)
	FromJSON(json.RawMessage) error
}
