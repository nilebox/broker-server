package api

import "encoding/json"

// Schemas represents a broker's schemas for both service instances and service
// bindings
type Schemas struct {
	Instance InstanceSchema `json:"service_instance"`
	Binding  BindingSchema  `json:"service_binding"`
}

type InstanceSchema struct {
	Create Schema `json:"create"`
	Update Schema `json:"update"`
}

type BindingSchema struct {
	Create Schema `json:"create"`
}

// Schema consists of the schema for inputs and the schema for outputs.
// Schemas are in the form of JSON Schema v4 (http://json-schema.org/).
type Schema struct {
	Parameters json.RawMessage `json:"parameters"`
}
