package retry

import (
	"context"
	"time"

	"github.com/nilebox/broker-server/pkg/stateful/storage"
	"github.com/nilebox/broker-server/pkg/stateful/task"
)

type watchDog struct {
	storage      storage.StorageWithLease
	tasks        []task.BrokerTask
	initialDelay time.Duration
	sleepDelay   time.Duration
}

func NewWatchDog(storage storage.StorageWithLease) *watchDog {
	return &watchDog{
		storage:      storage,
		tasks:        make([]task.BrokerTask, 0, 10),
		initialDelay: time.Second * 60,
		sleepDelay:   time.Second * 60,
	}
}

func (d *watchDog) add(task task.BrokerTask) {
	// TODO make it thread safe
	d.tasks = append(d.tasks, task)
}

func (d *watchDog) Run(ctx context.Context) {
	time.Sleep(d.initialDelay)

	for {
		d.checkTasks()

		select {
		case <-ctx.Done():
			// Received cancellation
			return
		default:
			// Sleep and continue running the loop
			time.Sleep(d.sleepDelay)
		}
	}
}

func (d *watchDog) checkTasks() {
	tasksCopy := make([]task.BrokerTask, 0, len(d.tasks))
	copy(tasksCopy, d.tasks)
	if len(tasksCopy) == 0 {
		return
	}

	runningTasks := make([]task.BrokerTask, 0, len(tasksCopy))
	finishedTasks := make([]task.BrokerTask, 0, len(tasksCopy))
	for _, t := range tasksCopy {
		switch t.State() {
		case task.BrokerTaskStateIdle:
			runningTasks = append(runningTasks, t)
		case task.BrokerTaskStateRunning:
			runningTasks = append(runningTasks, t)
		case task.BrokerTaskStateFinished:
			finishedTasks = append(runningTasks, t)
		default:
			panic("Unexpected task state: " + t.State())
		}
	}
	if len(runningTasks) > 0 {
		d.extendLease(runningTasks)
	}
	if len(finishedTasks) > 0 {
		d.tasks = runningTasks
	}
}
func (d *watchDog) extendLease(tasks []task.BrokerTask) error {
	instanceIds := make([]string, len(tasks))
	for i, t := range tasks {
		instanceIds[i] = t.InstanceId()
	}
	return d.storage.ExtendLease(instanceIds)
}
