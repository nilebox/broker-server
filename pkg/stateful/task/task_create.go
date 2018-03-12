package task

import "github.com/nilebox/broker-server/pkg/stateful/storage"

type CreateInstanceTask struct {
	brokerTask
	instance *storage.InstanceRecord
	storage  storage.Storage
	broker   Broker
}

func NewCreateTask(instance *storage.InstanceRecord, storage storage.Storage, broker Broker) *CreateInstanceTask {
	task := CreateInstanceTask{
		instance: instance,
		storage:  storage,
		broker:   broker,
	}
	t := brokerTask{
		State:   BrokerTaskStateIdle,
		RunFunc: task.run,
	}
	task.brokerTask = t
	return &task
}

func (t *CreateInstanceTask) run() {
	output, err := t.broker.CreateInstance(t.instance.InstanceId, t.instance.Parameters)
	if err != nil {
		t.storage.UpdateInstanceState(t.instance.InstanceId, storage.InstanceStateCreateFailed, err.Error())
		return
	}
	// TODO: persist outputs
	_ = output
	t.storage.UpdateInstanceState(t.instance.InstanceId, storage.InstanceStateCreateSucceeded, "")
}
