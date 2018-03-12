package task

type BrokerTask interface {
	Run()
}

type BrokerTaskRunner struct {
	State   BrokerTaskState
	RunFunc func()
}

func (t *BrokerTaskRunner) Run() {
	t.State = BrokerTaskStateRunning
	t.RunFunc()
	t.State = BrokerTaskStateFinished
}

type BrokerTaskState string

const (
	BrokerTaskStateIdle     BrokerTaskState = ""
	BrokerTaskStateRunning  BrokerTaskState = "Running"
	BrokerTaskStateFinished BrokerTaskState = "Finished"
)
