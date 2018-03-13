package task

import (
	"encoding/json"
	"github.com/nilebox/broker-server/pkg/stateful/storage"
)

// submitterStorageDecorator is a decorator around the storage.Storage
// interface, which submits the broker task every time the instance gets
// created or updated
type submitterStorageDecorator struct {
	storage     storage.Storage
	submitter   Submitter
	taskCreator *TaskCreator
}

func NewSubmitterStorageDecorator(storage storage.Storage, submitter Submitter, taskCreator *TaskCreator) storage.Storage {
	return &submitterStorageDecorator{
		storage:     storage,
		submitter:   submitter,
		taskCreator: taskCreator,
	}
}

func (d *submitterStorageDecorator) CreateInstance(instance *storage.InstanceRecord) error {
	err := d.storage.CreateInstance(instance)
	if err != nil {
		return err
	}
	return d.submitInstance(instance)
}

func (d *submitterStorageDecorator) UpdateInstance(instanceId string, parameters json.RawMessage, state storage.InstanceState) error {
	err := d.storage.UpdateInstance(instanceId, parameters, state)
	if err != nil {
		return err
	}
	return d.submitInstanceId(instanceId)
}

func (d *submitterStorageDecorator) UpdateInstanceState(instanceId string, state storage.InstanceState, e string) error {
	err := d.storage.UpdateInstanceState(instanceId, state, e)
	if err != nil {
		return err
	}
	return d.submitInstanceId(instanceId)
}

func (d *submitterStorageDecorator) GetInstance(instanceId string) (*storage.InstanceRecord, error) {
	return d.storage.GetInstance(instanceId)
}

func (d *submitterStorageDecorator) submitInstance(instance *storage.InstanceRecord) error {
	t, err := d.taskCreator.CreateTaskFor(instance)
	if err != nil {
		return err
	}
	d.submitter.Submit(t)
	return nil
}

func (d *submitterStorageDecorator) submitInstanceId(instanceId string) error {
	instance, err := d.storage.GetInstance(instanceId)
	if err != nil {
		return err
	}
	return d.submitInstance(instance)
}
