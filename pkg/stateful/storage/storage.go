package storage

import "encoding/json"

type Storage interface {
	CreateInstance(instance *InstanceRecord) error
	UpdateInstance(instanceId string, parameters json.RawMessage, state InstanceState) error
	UpdateInstanceState(instanceId string, state InstanceState, err string) error
	GetInstance(instanceId string) (*InstanceRecord, error)
	// TODO add methods for storing "instance outputs"
	// TODO add storage methods for bindings
}
