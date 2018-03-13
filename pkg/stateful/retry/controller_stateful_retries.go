package retry

import (
	"context"
	"github.com/nilebox/broker-server/pkg/stateful/task"
)

type retryController struct {
	taskExecutor  *taskExecutor
	taskScheduler *taskScheduler
	watchDog      *watchDog
}

func NewRetryController(storage StorageWithLease, broker task.Broker) *retryController {
	watchDog := NewWatchDog(storage)
	taskExecutor := NewTaskExecutor(watchDog)
	taskCreator := task.NewTaskCreator(storage, broker)
	taskScheduler := NewTaskScheduler(storage, taskExecutor, taskCreator)
	return &retryController{
		taskExecutor:  taskExecutor,
		taskScheduler: taskScheduler,
		watchDog:      watchDog,
	}
}

func (c *retryController) Start(ctx context.Context) {
	go c.taskScheduler.Run(ctx)
	go c.watchDog.Run(ctx)
}

func (c *retryController) Submit(task task.BrokerTask) {
	c.taskExecutor.Submit(task)
}
