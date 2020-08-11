package eagrors

import (
	"fmt"

	"github.com/ahmetcanozcan/eago/common/loggers"
)

// ProcessExitError invokes when error occured on JS process or
// process.exit function executed
type ProcessExitError struct {
	code int
}

// NewProcessExitError :
func NewProcessExitError(code int) error {
	return &ProcessExitError{code}
}

func (p ProcessExitError) Error() string {
	return fmt.Sprintf("Process exited with code %d", p.code)
}

// RecoverRuntime :
func RecoverRuntime(method, url string) {
	if p := recover(); p != nil {
		switch t := p.(type) {
		case error:
			e, ok := t.(*ProcessExitError)
			args := []interface{}{t, " in ", method, " ", url}
			if ok && e.code == 0 {
				loggers.Default().Info(args...)
			} else {
				loggers.Default().Error(args...)
			}
		default:
			panic(t)
		}

	}

}
