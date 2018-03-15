package retry

import "github.com/nilebox/broker-server/pkg/stateful/task"

type taskExecutor struct {
	watchDog *watchDog
}

func NewTaskExecutor(watchDog *watchDog) *taskExecutor {
	return &taskExecutor{
		watchDog: watchDog,
	}
}

func (e *taskExecutor) Submit(task task.BrokerTask) {
	e.watchDog.add(task)
	// TODO have a limited queue (pool) instead?
	go func() {
		err := task.Run()
		// TODO log err
		_ = err
	}()
}
