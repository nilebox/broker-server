package storage

import "fmt"

const (
	ErrorDeleted    = "instance does not exist or has been deleted"
	ErrorDeleting   = "instance is deleting"
	ErrorInProgress = "instance cannot be modified because it is in progress"
	ErrorConflict   = "instance already exists"

	ErrorUnknown = "unknown error"
)

type Storage interface {
	CreateInstance(instance *InstanceSpec) error
	UpdateInstance(instance *InstanceSpec) error
	DeleteInstance(instanceId string) error
	GetInstance(instanceId string) (*InstanceRecord, error)
	// TODO add methods for storing "instance outputs"
	// TODO add storage methods for bindings
}

type storageError struct {
	Err         string
	Description string
}

func NewDeleted(reason string, v ...interface{}) *storageError {
	return &storageError{
		Err:         ErrorDeleted,
		Description: fmt.Sprintf(reason, v...),
	}
}

func NewConflict(reason string, v ...interface{}) *storageError {
	return &storageError{
		Err:         ErrorConflict,
		Description: fmt.Sprintf(reason, v...),
	}
}

func NewDeleting(reason string, v ...interface{}) *storageError {
	return &storageError{
		Err:         ErrorDeleting,
		Description: fmt.Sprintf(reason, v...),
	}
}

func NewInProgress(reason string, v ...interface{}) *storageError {
	return &storageError{
		Err:         ErrorInProgress,
		Description: fmt.Sprintf(reason, v...),
	}
}

func (s *storageError) Error() string {
	return s.Err + ": " + s.Description
}

func IsDeletedError(e error) bool {
	return ReasonForError(e) == ErrorDeleted
}

func IsDeletingError(e error) bool {
	return ReasonForError(e) == ErrorDeleting
}

func IsInProgressError(e error) bool {
	return ReasonForError(e) == ErrorInProgress
}

func IsConflictError(e error) bool {
	return ReasonForError(e) == ErrorConflict
}

func ReasonForError(e error) string {
	switch t := e.(type) {
	case *storageError:
		return t.Err
	}
	return ErrorUnknown
}
