package task

import "github.com/nilebox/broker-server/pkg/stateful/storage"

type CreateInstanceTask struct {
	instanceId string
	storage    storage.Storage
	broker     Broker
}

func NewCreateTask(instanceId string, storage storage.Storage, broker Broker) BrokerTask {
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

func (t *CreateInstanceTask) run() {
	instance, err := t.storage.GetInstance(t.instanceId)
	if err != nil {
		// TODO Log error
		return
	}
	state, output, err := t.broker.CreateInstance(t.instanceId, instance.Parameters)
	if err != nil || state == ExecutionStateFailed {
		// TODO 'err' could mean a temporary error
		// Shall we have a separate error message for 'Failed' state?
		errorMessage := ""
		if err != nil {
			errorMessage = err.Error()
		}
		t.storage.UpdateInstanceState(t.instanceId, storage.InstanceStateCreateFailed, errorMessage)
		return
	}
	if state == ExecutionStateSuccess {
		instance.Outputs = output
		instance.State = storage.InstanceStateCreateSucceeded
		err = t.storage.UpdateInstance(&instance.InstanceSpec)
		if err != nil {
			// TODO Log error
			return
		}
		err = t.storage.UpdateInstanceState(t.instanceId, storage.InstanceStateCreateSucceeded, "")
		if err != nil {
			// TODO Log error
			return
		}
		return
	}
	// If InProgress - nothing to do
}
