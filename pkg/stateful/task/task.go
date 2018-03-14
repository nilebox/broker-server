package task

type BrokerTask interface {
	InstanceId() string
	State() BrokerTaskState
	Run()
}

type BrokerTaskRunner struct {
	instanceId string
	state      BrokerTaskState
	RunFunc    func()
}

func (t *BrokerTaskRunner) InstanceId() string {
	return t.instanceId
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
