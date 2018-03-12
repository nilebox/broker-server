package api

import "encoding/json"

// Instance represents an instance of a service
type Instance struct {
	ID               string `json:"id"`
	DashboardURL     string `json:"dashboard_url"`
	InternalID       string `json:"internal_id, omitempty"`
	ServiceID        string `json:"service_id"`
	PlanID           string `json:"plan_id"`
	OrganizationGUID string `json:"organization_guid"`
	SpaceGUID        string `json:"space_guid"`

	LastOperation *LastOperationResponse `json:"last_operation, omitempty"`

	Parameters map[string]interface{} `json:"parameters, omitempty"`
}

// CreateInstanceRequest represents a request to a broker to provision an
// instance of a service
type CreateInstanceRequest struct {
	OrgID      string          `json:"organization_guid"`
	PlanID     string          `json:"plan_id"`
	ServiceID  string          `json:"service_id"`
	SpaceID    string          `json:"space_guid"`
	Parameters json.RawMessage `json:"parameters,omitempty"`
}

// CreateInstanceResponse represents the response from a broker after a
// request to provision an instance of a service
type CreateInstanceResponse struct {
	DashboardURL string `json:"dashboard_url, omitempty"`
	Operation    string `json:"operation, omitempty"`
	Async        bool   `json:"-"`
}

// UpdateInstancePreviousValues represents a request to a broker containing
// the previous values in a UpdateInstanceRequest object.
type UpdateInstancePreviousValues struct {
	PlanID    *string `json:"plan_id,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
}

// UpdateInstanceRequest represents a request to a broker to update a
// instance of a service
type UpdateInstanceRequest struct {
	ServiceID      string                        `json:"service_id"`
	PlanID         *string                       `json:"plan_id,omitempty"`
	Parameters     json.RawMessage               `json:"parameters,omitempty"`
	PreviousValues *UpdateInstancePreviousValues `json:"previous_values,omitempty"`
}

// UpdateInstanceResponse represents the response from a broker after a
// request to update an instance of a service
type UpdateInstanceResponse struct {
	Operation string `json:"operation, omitempty"`
	Async     bool   `json:"-"`
}

// GetInstanceStatusResponse represents the response from a broker with a status
// of last operation applied to an instance of a service
type GetInstanceStatusResponse struct {
	State       string `json:"state, omitempty"`
	Description string `json:"description, omitempty"`
}

// DeleteInstanceResponse represents the response from a broker after a request
// to deprovision an instance of a service
type DeleteInstanceResponse struct {
	Operation string `json:"operation,omitempty"`
	Async     bool   `json:"-"`
}

// LastOperationResponse represents the broker response with the state of a discrete action
// that the broker is completing asynchronously
type LastOperationResponse struct {
	State       string `json:"state"`
	Description string `json:"description,omitempty"`
}

// Defines the possible states of an asynchronous request to a broker
const (
	StateInProgress = "in progress"
	StateSucceeded  = "succeeded"
	StateFailed     = "failed"
)
