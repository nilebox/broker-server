package api

import "encoding/json"

// Binding represents a binding to a service instance
type Binding struct {
	ID                string                 `json:"id"`
	ServiceID         string                 `json:"service_id"`
	AppID             string                 `json:"app_id"`
	ServicePlanID     string                 `json:"service_plan_id"`
	PrivateKey        string                 `json:"private_key"`
	ServiceInstanceID string                 `json:"service_instance_id"`
	BindResource      map[string]interface{} `json:"bind_resource,omitempty"`
	Parameters        json.RawMessage        `json:"parameters, omitempty"`
}

// BindingRequest represents a request to bind to a service instance
type BindingRequest struct {
	AppGUID      *string                `json:"app_guid,omitempty"`
	PlanID       string                 `json:"plan_id"`
	ServiceID    string                 `json:"service_id"`
	BindResource map[string]interface{} `json:"bind_resource,omitempty"`
	Parameters   json.RawMessage        `json:"parameters,omitempty"`
}

// CreateBindingResponse represents a response to a service binding
// request
type CreateBindingResponse struct {
	// SyslogDrainURL string      `json:"syslog_drain_url, omitempty"`
	Credentials json.RawMessage `json:"credentials"`
}
