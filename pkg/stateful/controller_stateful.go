package stateful

import (
	"context"

	"github.com/nilebox/broker-server/pkg/api"
	"github.com/nilebox/broker-server/pkg/controller"
	"github.com/nilebox/broker-server/pkg/stateful/storage"
	"github.com/nilebox/broker-server/pkg/zappers"
	"go.uber.org/zap"
)

type statefulController struct {
	appContext context.Context
	catalog    *api.Catalog
	storage    storage.Storage
}

func NewStatefulController(appContext context.Context, catalog *api.Catalog, storage storage.Storage) controller.BrokerController {
	return &statefulController{
		appContext: appContext,
		catalog:    catalog,
		storage:    storage,
	}
}

func (c *statefulController) Catalog(ctx context.Context) (*api.Catalog, error) {
	log := c.logger()
	log.Info("Catalog called")

	return c.catalog, nil
}

func (c *statefulController) GetInstanceStatus(ctx context.Context, instanceID, serviceID, planID, operation string) (*api.GetInstanceStatusResponse, error) {
	log := c.logger()
	log = log.With(zappers.InstanceID(instanceID))
	log.Info("GetInstanceStatus called")

	instanceRecord, err := c.storage.GetInstance(instanceID)
	if err != nil {
		return nil, err
	}
	instanceStateDescription, err := storage.GetInstanceStateDescription(instanceRecord.State)
	if err != nil {
		return nil, err
	}

	return &api.GetInstanceStatusResponse{
		State:       string(storage.GetOperationState(instanceRecord.State)),
		Description: string(instanceStateDescription),
	}, nil
}

func (c *statefulController) CreateInstance(ctx context.Context, instanceID string, acceptsIncomplete bool, req *api.CreateInstanceRequest) (*api.CreateInstanceResponse, error) {
	log := c.logger()
	log = log.With(zappers.InstanceID(instanceID))
	log.Info("CreateInstance called")

	// TODO check if instance already exists first
	//instanceRecord, err := c.storage.GetInstance(instanceID)
	//if err != nil {
	//	// TODO check for NotFound
	//	return nil, err
	//}
	//if instanceRecord != nil {
	//	// TODO return 409
	//}

	instanceParameters := &storage.InstanceSpec{
		InstanceId: instanceID,
		ServiceId:  req.ServiceID,
		PlanId:     req.PlanID,
		Parameters: req.Parameters,
	}
	err := c.storage.CreateInstance(instanceParameters)
	if err != nil {
		return nil, err
	}
	return &api.CreateInstanceResponse{
		Async: true,
	}, nil
}

func (c *statefulController) UpdateInstance(ctx context.Context, instanceID string, acceptsIncomplete bool, req *api.UpdateInstanceRequest) (*api.UpdateInstanceResponse, error) {
	log := c.logger()
	log = log.With(zappers.InstanceID(instanceID))
	log.Info("UpdateInstance called")

	instance, err := c.storage.GetInstance(instanceID)
	if err != nil {
		return nil, err
	}
	// TODO check for instance status first (should not have operations in progress)
	// instance.State = storage.InstanceStateUpdateInProgress
	if req.PlanID != nil {
		instance.PlanId = *req.PlanID
	}
	if req.Parameters != nil {
		instance.Parameters = req.Parameters
	}

	// Discard the state stuff
	err = c.storage.UpdateInstance(&instance.InstanceSpec)
	if err != nil {
		return nil, err
	}

	return &api.UpdateInstanceResponse{
		Async: true,
	}, nil
}

func (c *statefulController) RemoveInstance(ctx context.Context, instanceID, serviceID, planID string, acceptsIncomplete bool) (*api.DeleteInstanceResponse, error) {
	log := c.logger()
	log = log.With(zappers.InstanceID(instanceID))
	log.Info("RemoveInstance called")

	_, err := c.storage.GetInstance(instanceID)
	if err != nil {
		return nil, err
	}
	err = c.storage.DeleteInstance(instanceID)
	if err != nil {
		return nil, err
	}

	return &api.DeleteInstanceResponse{
		Async: true,
	}, nil
}

func (c *statefulController) CreateBinding(ctx context.Context, instanceID, bindingID string, req *api.BindingRequest) (*api.CreateBindingResponse, error) {
	log := c.logger()
	log = log.With(zappers.InstanceID(instanceID), zappers.BindingID(bindingID))
	log.Info("CreateBinding called")

	// TODO
	return nil, nil
}

func (c *statefulController) RemoveBinding(ctx context.Context, instanceID, bindingID, serviceID, planID string) error {
	log := c.logger()
	log = log.With(zappers.InstanceID(instanceID), zappers.BindingID(bindingID))
	log.Info("RemoveBinding called")

	// TODO
	return nil
}

func (c *statefulController) logger() *zap.Logger {
	return c.appContext.Value("log").(*zap.Logger)
}
