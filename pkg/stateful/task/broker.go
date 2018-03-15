package task

import "encoding/json"

type Broker interface {
	CreateInstance(instanceId string, parameters json.RawMessage) (json.RawMessage, error)
	UpdateInstance(instanceId string, parameters json.RawMessage) (json.RawMessage, error)
	DeleteInstance(instanceId string, parameters json.RawMessage) error
	//CreateBinding(instanceId string, instanceParameters json.RawMessage, instanceOutputs json.RawMessage,
	//	bindingId string, bindindParameters json.RawMessage) (ExecutionState, json.RawMessage, error)
	//DeleteBinding(instanceId string, instanceParameters json.RawMessage, instanceOutputs json.RawMessage,
	//	bindingId string, bindindParameters json.RawMessage) (ExecutionState, error)
}

type ExecutionState string

const (
	ExecutionStateSuccess    ExecutionState = "Success"
	ExecutionStateFailed     ExecutionState = "Failed"
	ExecutionStateInProgress ExecutionState = "InProgress"
)
