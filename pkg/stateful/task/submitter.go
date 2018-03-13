package task

type Submitter interface {
	Submit(task BrokerTask)
}
