package task

import "github.com/nilebox/broker-server/pkg/stateful/storage"

var _ storage.Storage = &submitterStorageDecorator{}
