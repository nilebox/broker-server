package storage

import (
	"encoding/json"
	"errors"
	"github.com/nilebox/broker-server/pkg/api"
)

type InstanceRecord struct {
	InstanceId string
	// ServiceId + PlanId?
	// resourceType ResourceType
	Parameters json.RawMessage
	State      InstanceState
	Error      string
}

type InstanceState string

const (
	InstanceStateCreateInProgress InstanceState = "CreateInProgress"
	InstanceStateCreateSucceeded  InstanceState = "CreateSucceeded"
	InstanceStateCreateFailed     InstanceState = "CreateFailed"
	InstanceStateUpdateInProgress InstanceState = "UpdateInProgress"
	InstanceStateUpdateSucceeded  InstanceState = "UpdateSucceeded"
	InstanceStateUpdateFailed     InstanceState = "UpdateFailed"
	InstanceStateDeleteInProgress InstanceState = "DeleteInProgress"
	InstanceStateDeleteSucceeded  InstanceState = "DeleteSucceeded"
	InstanceStateDeleteFailed     InstanceState = "DeleteFailed"
)

type InstanceStateDescription string

const (
	InstanceStateDescriptionCreateInProgress InstanceStateDescription = "Instance is being provisioned"
	InstanceStateDescriptionCreateSucceeded  InstanceStateDescription = "Instance has been provisioned"
	InstanceStateDescriptionCreateFailed     InstanceStateDescription = "Failed to provision an instance "
	InstanceStateDescriptionUpdateInProgress InstanceStateDescription = "Instance is being updated"
	InstanceStateDescriptionUpdateSucceeded  InstanceStateDescription = "Instance has been updated"
	InstanceStateDescriptionUpdateFailed     InstanceStateDescription = "Failed to update an instance"
	InstanceStateDescriptionDeleteInProgress InstanceStateDescription = "Instance is being deprovisioned"
	InstanceStateDescriptionDeleteSucceeded  InstanceStateDescription = "Instance has been deprovisioned"
	InstanceStateDescriptionDeleteFailed     InstanceStateDescription = "Failed to deprovision an instance"
)

func GetInstanceStateDescription(state InstanceState) (InstanceStateDescription, error) {
	switch state {
	case InstanceStateCreateInProgress:
		return InstanceStateDescriptionCreateInProgress, nil
	// TODO: add remaining state descriptions
	default:
		return "", errors.New("Unexpected instance state: " + state)

	}
}

func IsInProgress(state InstanceState) bool {
	switch state {
	case InstanceStateCreateInProgress:
		return true
	case InstanceStateUpdateInProgress:
		return true
	case InstanceStateDeleteInProgress:
		return true
	default:
		return false
	}
}

func GetOperationState(state InstanceState) api.OperationState {
	switch state {
	case InstanceStateCreateInProgress:
		return api.OperationStateInProgress
	// TODO add the rest InstanceState values
	default:
		panic("Unexpected state: " + state)
	}
}
