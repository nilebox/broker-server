package task

import (
	"errors"
	"github.com/nilebox/broker-server/pkg/stateful/storage"
)

type UpdateInstanceTask struct {
	instanceId string
	storage    storage.StorageWithLease
	broker     Broker
}

func NewUpdateTask(instanceId string, storage storage.StorageWithLease, broker Broker) BrokerTask {
	task := UpdateInstanceTask{
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

func (t *UpdateInstanceTask) run() error {
	instance, err := t.storage.GetInstance(t.instanceId)
	if err != nil {
		return err
	}
	if instance.State != storage.InstanceStateUpdateInProgress {
		return errors.New("Unexpected status: " + string(instance.State))
	}
	output, err := t.broker.UpdateInstance(t.instanceId, instance.Spec.Parameters)
	if err != nil {
		// TODO 'err' could mean a temporary error
		// Shall we have a separate error message for 'Failed' state?
		errorMessage := ""
		if err != nil {
			errorMessage = err.Error()
		}
		t.storage.UpdateInstanceState(t.instanceId, storage.InstanceStateUpdateFailed, errorMessage)
		return err
	}

	err = t.storage.UpdateInstanceOutputs(t.instanceId, output)
	if err != nil {
		return err
	}
	err = t.storage.UpdateInstanceState(t.instanceId, storage.InstanceStateUpdateSucceeded, "")
	if err != nil {
		return err
	}
	return nil
}
