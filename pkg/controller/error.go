package controller

import (
	"fmt"
	"net/http"
)

type controllerError struct {
	Code        int
	Err         string
	Description string
}

func NewConflict(reason string, v ...interface{}) error {
	return &controllerError{
		Code: http.StatusConflict,
		Err:  fmt.Sprintf(reason, v...),
	}
}

func NewGone(reason string, v ...interface{}) error {
	return &controllerError{
		Code: http.StatusGone,
		Err:  fmt.Sprintf(reason, v...),
	}
}

func NewAsyncRequired(reason string, v ...interface{}) error {
	return NewUnprocessableEntity("This service plan requires client support for asynchronous service operations.")
}

func NewUnprocessableEntity(reason string, v ...interface{}) error {
	return &controllerError{
		Code: http.StatusUnprocessableEntity,
		Err:  fmt.Sprintf(reason, v...),
	}
}

func NewBadRequest(reason string, v ...interface{}) error {
	return &controllerError{
		Code: http.StatusBadRequest,
		Err:  fmt.Sprintf(reason, v...),
	}
}

func NewInternalServerError(reason string, v ...interface{}) error {
	return &controllerError{
		Code: http.StatusInternalServerError,
		Err:  fmt.Sprintf(reason, v...),
	}
}

func (c *controllerError) Error() string {
	return fmt.Sprintf("%v %v: %v %v", c.Code, http.StatusText(c.Code), c.Err, c.Description)
}

func GetControllerError(e error) *controllerError {
	switch t := e.(type) {
	case *controllerError:
		return t
	}
	return nil
}
