package controller

import (
	"context"

	"github.com/nilebox/broker-server/pkg/api"
)

// Controller defines the APIs that all controllers are expected to support. Implementations
// should be concurrency-safe
type BrokerController interface {
	Catalog(ctx context.Context) (*api.Catalog, error)

	GetInstanceStatus(ctx context.Context, instanceID, serviceID, planID, operation string) (*api.GetServiceInstanceStatusResponse, error)
	CreateInstance(ctx context.Context, instanceID string, acceptsIncomplete bool, req *api.CreateServiceInstanceRequest) (*api.CreateServiceInstanceResponse, error)
	UpdateInstance(ctx context.Context, instanceID string, acceptsIncomplete bool, req *api.UpdateServiceInstanceRequest) (*api.UpdateServiceInstanceResponse, error)
	RemoveInstance(ctx context.Context, instanceID, serviceID, planID string, acceptsIncomplete bool) (*api.DeleteServiceInstanceResponse, error)

	CreateBinding(ctx context.Context, instanceID, bindingID string, req *api.BindingRequest) (*api.CreateServiceBindingResponse, error)
	RemoveBinding(ctx context.Context, instanceID, bindingID, serviceID, planID string) error
}
