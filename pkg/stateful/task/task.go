package task

type brokerTask struct {
	State   BrokerTaskState
	RunFunc func()
}

func (t *brokerTask) run() {
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
