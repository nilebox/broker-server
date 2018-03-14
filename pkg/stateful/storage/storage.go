package storage

type Storage interface {
	CreateInstance(instance *InstanceRecord) error
	UpdateInstance(instance *InstanceRecord) error
	UpdateInstanceState(instanceId string, state InstanceState, err string) error
	GetInstance(instanceId string) (*InstanceRecord, error)
	// TODO add methods for storing "instance outputs"
	// TODO add storage methods for bindings
}
