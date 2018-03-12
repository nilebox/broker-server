package task

import "encoding/json"

type Broker interface {
	CreateInstance(instanceId string, parameters json.RawMessage) (json.RawMessage, error)
	UpdateInstance(instanceId string, parameters json.RawMessage) (json.RawMessage, error)
	DeleteInstance(instanceId string, parameters json.RawMessage) error
	// TODO add methods for bindings
}
