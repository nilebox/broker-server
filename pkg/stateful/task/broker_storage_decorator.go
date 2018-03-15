package task

import (
	"encoding/json"
	"github.com/nilebox/broker-server/pkg/stateful/retry"
	"github.com/nilebox/broker-server/pkg/stateful/storage"
)

// brokerStorageDecorator is a decorator around the Broker
// interface, which automatically submits state updates to the storage
type brokerStorageDecorator struct {
	storage retry.StorageWithLease
	broker  Broker
}

func (d *brokerStorageDecorator) CreateInstance(instanceId string, parameters json.RawMessage) (ExecutionState, json.RawMessage, error) {
	instance, err := d.storage.GetInstance(instanceId)
	if err != nil {
		return "", nil, err
	}
	d.storage.UpdateInstanceState(instanceId, storage.InstanceStateCreateInProgress, "")
	state, output, err := d.broker.CreateInstance(instanceId, parameters)
	if err != nil || state == ExecutionStateFailed {
		// TODO 'err' could mean a temporary error
		// Shall we have a separate error message for 'Failed' state?
		errorMessage := ""
		if err != nil {
			errorMessage = err.Error()
		}
		d.storage.UpdateInstanceState(instanceId, storage.InstanceStateCreateFailed, errorMessage)
		return state, output, err
	}
	if state == ExecutionStateSuccess {
		instance.Spec.Outputs = output
		err = d.storage.UpdateInstance(&instance.Spec)
		if err != nil {
			// TODO Log error
			return "", nil, err
		}
		err = d.storage.UpdateInstanceState(instanceId, storage.InstanceStateCreateSucceeded, "")
		if err != nil {
			// TODO Log error
			return "", nil, err
		}
	}
	// InProgress - nothing to do

	return state, output, err
}

func (d *brokerStorageDecorator) UpdateInstance(instanceId string, parameters json.RawMessage) (ExecutionState, json.RawMessage, error) {
	return "", nil, nil
}

func (d *brokerStorageDecorator) DeleteInstance(instanceId string, parameters json.RawMessage) (ExecutionState, error) {
	return "", nil
}
