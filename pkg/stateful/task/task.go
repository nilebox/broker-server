package task

type BrokerTask interface {
	InstanceId() string
	State() BrokerTaskState
	Run() error
}

type BrokerTaskRunner struct {
	instanceId string
	state      BrokerTaskState
	RunFunc    func() error
}

func (t *BrokerTaskRunner) InstanceId() string {
	return t.instanceId
}

func (t *BrokerTaskRunner) State() BrokerTaskState {
	return t.state
}

func (t *BrokerTaskRunner) Run() error {
	defer func() {
		t.state = BrokerTaskStateFinished
	}()

	t.state = BrokerTaskStateRunning
	err := t.RunFunc()
	return err
}

type BrokerTaskState string

const (
	BrokerTaskStateIdle     BrokerTaskState = ""
	BrokerTaskStateRunning  BrokerTaskState = "Running"
	BrokerTaskStateFinished BrokerTaskState = "Finished"
)
