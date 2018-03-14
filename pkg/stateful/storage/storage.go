package storage

import "fmt"

const (
	errorDeleted    = "instance does not exist or has been deleted"
	errorDeleting   = "instance is deleting"
	errorInProgress = "instance cannot be modified because it is in progress"
	errorConflict   = "instance already exists"

	errorUnknown = "unknown error"
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
		Err:         errorDeleted,
		Description: fmt.Sprintf(reason, v...),
	}
}

func NewConflict(reason string, v ...interface{}) *storageError {
	return &storageError{
		Err:         errorConflict,
		Description: fmt.Sprintf(reason, v...),
	}
}

func NewDeleting(reason string, v ...interface{}) *storageError {
	return &storageError{
		Err:         errorDeleting,
		Description: fmt.Sprintf(reason, v...),
	}
}

func NewInProgress(reason string, v ...interface{}) *storageError {
	return &storageError{
		Err:         errorInProgress,
		Description: fmt.Sprintf(reason, v...),
	}
}

func (s *storageError) Error() string {
	return s.Err + ": " + s.Description
}

func IsDeletedError(e error) bool {
	return ReasonForError(e) == errorDeleted
}

func IsDeletingError(e error) bool {
	return ReasonForError(e) == errorDeleting
}

func IsInProgressError(e error) bool {
	return ReasonForError(e) == errorInProgress
}

func IsConflictError(e error) bool {
	return ReasonForError(e) == errorConflict
}

func ReasonForError(e error) string {
	switch t := e.(type) {
	case *storageError:
		return t.Err
	}
	return errorUnknown
}
