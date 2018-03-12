package stateful

import "github.com/nilebox/broker-server/pkg/controller"

var _ controller.BrokerController = &statefulController{}
