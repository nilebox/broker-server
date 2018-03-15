package task

import (
	"errors"
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
	if instance.State != storage.InstanceStateCreateInProgress {
		return errors.New("Unexpected status: " + string(instance.State))
	}
	output, err := t.broker.CreateInstance(t.instanceId, instance.Spec.Parameters)
	if err != nil {
		// TODO 'err' could mean a temporary error
		// Shall we have a separate error message for 'Failed' state?
		errorMessage := ""
		if err != nil {
			errorMessage = err.Error()
		}
		t.storage.UpdateInstanceState(t.instanceId, storage.InstanceStateCreateFailed, errorMessage)
		return err
	}

	err = t.storage.UpdateInstanceOutputs(t.instanceId, output)
	if err != nil {
		return err
	}
	err = t.storage.UpdateInstanceState(t.instanceId, storage.InstanceStateCreateSucceeded, "")
	if err != nil {
		return err
	}
	return nil
}
