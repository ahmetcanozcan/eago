package eagrors

import (
	"errors"

	"github.com/ahmetcanozcan/eago/common/loggers"
)

// StackError represents errors that refers another or referred from another error
type StackError interface {
	causer
	error
}

type stackError struct {
	cause error
	err   error
}

func (s *stackError) Error() string {
	return s.err.Error()
}

func (s *stackError) Cause() error {
	return s.cause
}

// NewError returns new stack error
func NewError(errStr string) error {
	err := errors.New(errStr)
	return &stackError{err: err, cause: nil}
}

// NewErrorWithCause creates a stack error with given error
func NewErrorWithCause(cause error, err string) error {
	return &stackError{cause, errors.New(err)}
}

// PrintError print errors and causes to default logger
func PrintError(err error) {
	FprintStackError(loggers.Default(), err)
}

// GetErrorString returns errors string of an error
// if it StackError then get all cause
func GetErrorString(err error) string {
	errMsg := ""
	for {
		if err == nil {
			break
		}
		switch e := err.(type) {
		case StackError:
			errMsg += e.Error() + "\n"
			err = e.Cause()
		default:
			errMsg += err.Error()
			return errMsg
		}
	}
	return errMsg
}

// FprintStackError print errors and causes to given error logger
func FprintStackError(l interface{ Error(...interface{}) }, err error) {
	for {
		if err == nil {
			return
		}
		switch s := err.(type) {
		case StackError:
			l.Error(s)
			err = s.Cause()
		default:
			l.Error(err)
			return
		}
	}
}

// HandleErrors return first error which is not nil
func HandleErrors(_errors ...error) error {
	for _, err := range _errors {
		if err != nil {
			return err
		}
	}
	return nil
}
