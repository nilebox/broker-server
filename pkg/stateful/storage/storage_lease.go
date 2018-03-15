package storage

type StorageWithLease interface {
	Storage
	UpdateInstanceState(instanceId string, state InstanceState, err string) error
	ExtendLease(instanceIds []string) error
	LeaseAbandonedInstances(maxBatchSize uint32) ([]*InstanceRecord, error)
}
