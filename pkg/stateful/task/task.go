package task

import "github.com/nilebox/broker-server/pkg/stateful/storage"

type BrokerTask interface {
	Instance() *storage.InstanceRecord
	State() BrokerTaskState
	Run()
}

type BrokerTaskRunner struct {
	instance *storage.InstanceRecord
	state    BrokerTaskState
	RunFunc  func()
}

func (t *BrokerTaskRunner) Instance() *storage.InstanceRecord {
	return t.instance
}

func (t *BrokerTaskRunner) State() BrokerTaskState {
	return t.state
}

func (t *BrokerTaskRunner) Run() {
	t.state = BrokerTaskStateRunning
	t.RunFunc()
	t.state = BrokerTaskStateFinished
}

type BrokerTaskState string

const (
	BrokerTaskStateIdle     BrokerTaskState = ""
	BrokerTaskStateRunning  BrokerTaskState = "Running"
	BrokerTaskStateFinished BrokerTaskState = "Finished"
)
