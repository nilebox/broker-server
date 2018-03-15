package retry

import "github.com/nilebox/broker-server/pkg/stateful/storage"

type StorageWithLease interface {
	storage.Storage
	UpdateInstanceState(instanceId string, state storage.InstanceState, err string) error
	ExtendLease(instanceIds []string) error
	LeaseAbandonedInstances(maxBatchSize uint32) ([]*storage.InstanceRecord, error)
}
