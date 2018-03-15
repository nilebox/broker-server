package task

import (
	"errors"

	"github.com/nilebox/broker-server/pkg/stateful/storage"
)

type TaskCreator struct {
	storage storage.StorageWithLease
	broker  Broker
}

func NewTaskCreator(storage storage.StorageWithLease, broker Broker) *TaskCreator {
	return &TaskCreator{
		storage: storage,
		broker:  broker,
	}
}

func (tc *TaskCreator) CreateTaskFor(instance *storage.InstanceRecord) (BrokerTask, error) {
	switch instance.State {
	case storage.InstanceStateCreateInProgress:
		return NewCreateTask(instance.Spec.InstanceId, tc.storage, tc.broker), nil
	case storage.InstanceStateUpdateInProgress:
		return NewUpdateTask(instance.Spec.InstanceId, tc.storage, tc.broker), nil
	case storage.InstanceStateDeleteInProgress:
		return NewDeleteTask(instance.Spec.InstanceId, tc.storage, tc.broker), nil
	default:
		// There is no operation in progress.
		return nil, errors.New("Instance is not in progress: " + string(instance.State))
	}
}
