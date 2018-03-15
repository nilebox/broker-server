package retry

import (
	"context"
	"time"

	"github.com/nilebox/broker-server/pkg/stateful/storage"
	"github.com/nilebox/broker-server/pkg/stateful/task"
)

const (
	maxBatchSize = 10
)

type taskScheduler struct {
	storage      storage.StorageWithLease
	taskExecutor *taskExecutor
	taskCreator  *task.TaskCreator
	initialDelay time.Duration
	sleepDelay   time.Duration
}

func NewTaskScheduler(storage storage.StorageWithLease, taskExecutor *taskExecutor, taskCreator *task.TaskCreator) *taskScheduler {
	return &taskScheduler{
		storage:      storage,
		taskExecutor: taskExecutor,
		taskCreator:  taskCreator,
		initialDelay: time.Second * 60,
		sleepDelay:   time.Second * 60,
	}
}

func (s *taskScheduler) Run(ctx context.Context) {
	time.Sleep(s.initialDelay)

	for {
		instances, err := s.storage.LeaseAbandonedInstances(maxBatchSize)
		if err != nil {
			// TODO log error
		}
		if err == nil && len(instances) != 0 {
			s.submitTasks(instances)
		}

		select {
		case <-ctx.Done():
			// Received cancellation
			return
		default:
			// Sleep and continue running the loop
			time.Sleep(s.sleepDelay)
		}
	}
}

func (s *taskScheduler) submitTasks(instances []*storage.InstanceRecord) error {
	for _, instance := range instances {
		t, err := s.taskCreator.CreateTaskFor(instance)
		if err != nil {
			return err
		}
		s.taskExecutor.Submit(t)
	}
	return nil
}
