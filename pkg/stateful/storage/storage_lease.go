package storage

import "encoding/json"

type StorageWithLease interface {
	Storage
	UpdateInstanceOutputs(instanceId string, outputs json.RawMessage) error
	UpdateInstanceState(instanceId string, state InstanceState, err string) error
	ExtendLease(instanceIds []string) error
	LeaseAbandonedInstances(maxBatchSize uint32) ([]*InstanceRecord, error)
}
