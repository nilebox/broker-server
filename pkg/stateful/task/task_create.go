package task

import "github.com/nilebox/broker-server/pkg/stateful/storage"

type CreateInstanceTask struct {
	instance *storage.InstanceRecord
	storage  storage.Storage
	broker   Broker
}

func NewCreateTask(instance *storage.InstanceRecord, storage storage.Storage, broker Broker) BrokerTask {
	task := CreateInstanceTask{
		instance: instance,
		storage:  storage,
		broker:   broker,
	}
	runner := BrokerTaskRunner{
		instance: instance,
		state:    BrokerTaskStateIdle,
		RunFunc:  task.run,
	}
	return &runner
}

func (t *CreateInstanceTask) run() {
	state, output, err := t.broker.CreateInstance(t.instance.InstanceId, t.instance.Parameters)
	if err != nil || state == ExecutionStateFailed {
		// TODO 'err' could mean a temporary error
		// Shall we have a separate error message for 'Failed' state?
		errorMessage := ""
		if err != nil {
			errorMessage = err.Error()
		}
		t.storage.UpdateInstanceState(t.instance.InstanceId, storage.InstanceStateCreateFailed, errorMessage)
		return
	}
	if state == ExecutionStateSuccess {
		// TODO: persist outputs (add method to storage)
		_ = output
		t.storage.UpdateInstanceState(t.instance.InstanceId, storage.InstanceStateCreateSucceeded, "")
		return
	}
	// If InProgress - nothing to do
}
