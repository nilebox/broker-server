package retry

import "github.com/nilebox/broker-server/pkg/stateful/storage"

type StorageWithLease interface {
	storage.Storage
	ExtendLease(instances []*storage.InstanceRecord) error
	LeaseAbandonedInstances(maxBatchSize uint32) []*storage.InstanceRecord
}
