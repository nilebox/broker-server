package task

import (
	"errors"
	"github.com/nilebox/broker-server/pkg/stateful/storage"
)

type DeleteInstanceTask struct {
	instanceId string
	storage    storage.StorageWithLease
	broker     Broker
}

func NewDeleteTask(instanceId string, storage storage.StorageWithLease, broker Broker) BrokerTask {
	task := DeleteInstanceTask{
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

func (t *DeleteInstanceTask) run() error {
	instance, err := t.storage.GetInstance(t.instanceId)
	if err != nil {
		return err
	}
	if instance.State != storage.InstanceStateDeleteInProgress {
		return errors.New("Unexpected status: " + string(instance.State))
	}
	err = t.broker.DeleteInstance(t.instanceId, instance.Spec.Parameters)
	if err != nil {
		// TODO 'err' could mean a temporary error
		// Shall we have a separate error message for 'Failed' state?
		errorMessage := ""
		if err != nil {
			errorMessage = err.Error()
		}
		t.storage.UpdateInstanceState(t.instanceId, storage.InstanceStateDeleteFailed, errorMessage)
		return err
	}

	err = t.storage.UpdateInstanceState(t.instanceId, storage.InstanceStateDeleteSucceeded, "")
	if err != nil {
		return err
	}
	return nil
}
