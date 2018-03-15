package retry

import (
	"context"

	"github.com/nilebox/broker-server/pkg/stateful/storage"
	"github.com/nilebox/broker-server/pkg/stateful/task"
)

type retryController struct {
	storage       storage.StorageWithLease
	taskCreator   *task.TaskCreator
	taskExecutor  *taskExecutor
	taskScheduler *taskScheduler
	watchDog      *watchDog
}

func NewRetryController(storage storage.StorageWithLease, broker task.Broker) *retryController {
	watchDog := NewWatchDog(storage)
	taskExecutor := NewTaskExecutor(watchDog)
	taskCreator := task.NewTaskCreator(storage, broker)
	taskScheduler := NewTaskScheduler(storage, taskExecutor, taskCreator)
	return &retryController{
		storage:       storage,
		taskCreator:   taskCreator,
		taskExecutor:  taskExecutor,
		taskScheduler: taskScheduler,
		watchDog:      watchDog,
	}
}

func (c *retryController) Start(ctx context.Context) {
	go c.taskScheduler.Run(ctx)
	go c.watchDog.Run(ctx)
}

// CreateStorageWithSubmitter returns a storage decorated with task submitter
func (c *retryController) CreateStorageWithSubmitter() storage.Storage {
	return task.NewSubmitterStorageDecorator(c.storage, c.taskExecutor, c.taskCreator)
}
