package task

import (
	"errors"

	"github.com/nilebox/broker-server/pkg/stateful/storage"
)

type TaskCreator struct {
	storage storage.Storage
	broker  Broker
}

func NewTaskCreator(storage storage.Storage, broker Broker) *TaskCreator {
	return &TaskCreator{
		storage: storage,
		broker:  broker,
	}
}

func (tc *TaskCreator) CreateTaskFor(instance *storage.InstanceRecord) (BrokerTask, error) {
	switch instance.State {
	case storage.InstanceStateCreateInProgress:
		return NewCreateTask(instance.InstanceId, tc.storage, tc.broker), nil
	// Add missing in progress states
	default:
		// There is no operation in progress.
		return nil, errors.New("Instance is not in progress: " + string(instance.State))
	}
}
