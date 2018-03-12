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
		State:   BrokerTaskStateIdle,
		RunFunc: task.run,
	}
	return &runner
}

func (t *CreateInstanceTask) run() {
	output, err := t.broker.CreateInstance(t.instance.InstanceId, t.instance.Parameters)
	if err != nil {
		t.storage.UpdateInstanceState(t.instance.InstanceId, storage.InstanceStateCreateFailed, err.Error())
		return
	}
	// TODO: persist outputs (add method to storage)
	_ = output
	t.storage.UpdateInstanceState(t.instance.InstanceId, storage.InstanceStateCreateSucceeded, "")
}
