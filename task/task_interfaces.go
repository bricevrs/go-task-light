package task

/*

The goal of this file is to define how every task should behave.
For instance, a task should be able to be added to the database, updated, and retrieved.

*/

import "encoding/json"

type TaskQuerier interface {
	Add() error
	UpdateStatus(status TaskStatus) error
	Execute() error
	Relaunch() error
}

type TaskPayloadMarshaler interface {
	ToJSON() ([]byte, error)
	FromJSON(json.RawMessage) error
}
