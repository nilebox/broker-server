package task

import (
	"github.com/nilebox/broker-server/pkg/stateful/storage"
)

type CreateInstanceTask struct {
	instanceId string
	storage    storage.StorageWithLease
	broker     Broker
}

func NewCreateTask(instanceId string, storage storage.StorageWithLease, broker Broker) BrokerTask {
	task := CreateInstanceTask{
		instanceId: instanceId,
		storage:    storage,
		broker:     broker,
	}
	runner := BrokerTaskRunner{
		instanceId: instanceId,
		state:      BrokerTaskStateIdle,
		RunFunc:    task.run,
	}
	return &runner
}

func (t *CreateInstanceTask) run() error {
	instance, err := t.storage.GetInstance(t.instanceId)
	if err != nil {
		return err
	}
	t.storage.UpdateInstanceState(t.instanceId, storage.InstanceStateCreateInProgress, "")
	state, output, err := t.broker.CreateInstance(t.instanceId, instance.Spec.Parameters)
	if err != nil || state == ExecutionStateFailed {
		// TODO 'err' could mean a temporary error
		// Shall we have a separate error message for 'Failed' state?
		errorMessage := ""
		if err != nil {
			errorMessage = err.Error()
		}
		t.storage.UpdateInstanceState(t.instanceId, storage.InstanceStateCreateFailed, errorMessage)
		return err
	}
	if state == ExecutionStateSuccess {
		instance.Spec.Outputs = output
		err = t.storage.UpdateInstance(&instance.Spec)
		if err != nil {
			return err
		}
		err = t.storage.UpdateInstanceState(t.instanceId, storage.InstanceStateCreateSucceeded, "")
		if err != nil {
			return err
		}
	}

	// InProgress - nothing to do
	return nil
}
