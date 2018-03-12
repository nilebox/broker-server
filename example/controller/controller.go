package controller

import (
	"context"

	"github.com/nilebox/broker-server/pkg/api"
	"github.com/nilebox/broker-server/pkg/controller"
	"github.com/nilebox/broker-server/pkg/zappers"
	"go.uber.org/zap"
)

const (
	// TODO nilebox: support multiple services and plans
	ServiceId   = "uuid1"
	ServiceName = "my-service"
	PlanId      = "uuid2"
	PlanName    = "default"

	StatusInProgress = "in progress"
	StatusSucceeded  = "succeeded"
	StatusFailed     = "failed"
)

type exampleController struct {
	appCtx context.Context
}

func NewController(ctx context.Context) (controller.BrokerController, error) {
	return &exampleController{
		appCtx: ctx,
	}, nil
}

func (c *exampleController) Catalog(ctx context.Context) (*api.Catalog, error) {
	log := c.appCtx.Value("log").(*zap.Logger)
	log.Info("Catalog called")

	catalog := api.Catalog{
		Services: []*api.Service{
			{
				ID:          ServiceId,
				Name:        ServiceName,
				Description: "Service description",
				Bindable:    true,
				Plans: []api.ServicePlan{
					{
						ID:          PlanId,
						Name:        PlanName,
						Description: "Plan description",
						//Schemas: &api.Schemas{
						//	Instance: api.InstanceSchema{
						//		Create: api.Schema{
						//			Parameters: c.schema,
						//		},
						//		Update: api.Schema{
						//			Parameters: c.schema,
						//		},
						//	},
						//},
					},
				},
				PlanUpdateable: false,
			},
		},
	}

	return &catalog, nil
}

func (c *exampleController) GetInstanceStatus(ctx context.Context, instanceID, serviceID, planID, operation string) (*api.GetInstanceStatusResponse, error) {
	log := ctx.Value("log").(*zap.Logger)
	log = log.With(zappers.InstanceID(instanceID))
	log.Info("GetInstanceStatus called")
	return nil, nil
}

func (c *exampleController) CreateInstance(ctx context.Context, instanceID string, acceptsIncomplete bool, req *api.CreateInstanceRequest) (*api.CreateInstanceResponse, error) {
	log := ctx.Value("log").(*zap.Logger)
	log = log.With(zappers.InstanceID(instanceID))
	log.Info("CreateInstance called")
	return nil, nil
}

func (c *exampleController) UpdateInstance(ctx context.Context, instanceID string, acceptsIncomplete bool, req *api.UpdateInstanceRequest) (*api.UpdateInstanceResponse, error) {
	log := ctx.Value("log").(*zap.Logger)
	log = log.With(zappers.InstanceID(instanceID))
	log.Info("UpdateInstance called")
	return nil, nil
}

func (c *exampleController) RemoveInstance(ctx context.Context, instanceID, serviceID, planID string, acceptsIncomplete bool) (*api.DeleteInstanceResponse, error) {
	log := ctx.Value("log").(*zap.Logger)
	log = log.With(zappers.InstanceID(instanceID))
	log.Info("RemoveInstance called")
	return nil, nil
}

func (c *exampleController) CreateBinding(ctx context.Context, instanceID, bindingID string, req *api.BindingRequest) (*api.CreateServiceBindingResponse, error) {
	log := ctx.Value("log").(*zap.Logger)
	log = log.With(zappers.InstanceID(instanceID), zappers.BindingID(bindingID))
	log.Info("CreateBinding called")
	return nil, nil
}

func (c *exampleController) RemoveBinding(ctx context.Context, instanceID, bindingID, serviceID, planID string) error {
	log := ctx.Value("log").(*zap.Logger)
	log = log.With(zappers.InstanceID(instanceID), zappers.BindingID(bindingID))
	log.Info("RemoveBinding called")
	return nil
}
