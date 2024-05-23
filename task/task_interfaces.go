package task

import "encoding/json"

type TaskQuerier interface {
	Add(arg ...interface{}) error
	UpdateStatus(arg ...interface{}) error
	Get(arg ...interface{}) (*Task, error)
	Execute() error
	Relaunch() error
}

type TaskPayloadMarshaler interface {
	ToJSON() ([]byte, error)
	FromJSON(json.RawMessage) error
}
