package eagrors

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

const (
	argumentCountErorHeader = "Unexpected argument count"
)

// ArgumentCountError invokes when otto function is passed
// by unexpected number of parameters
type ArgumentCountError struct {
	want int
	got  int
}

func (e ArgumentCountError) Error() string {
	return fmt.Sprintf("%s : %s",
		argumentCountErorHeader,
		getArgumentCountErrorBodyText(e.want, e.got))
}

// JSError return otto error
func (e ArgumentCountError) JSError(vm *otto.Otto) otto.Value {
	return vm.MakeCustomError(argumentCountErorHeader,
		getArgumentCountErrorBodyText(e.want, e.got))
}

func getArgumentCountErrorBodyText(want, got int) string {
	return fmt.Sprintf("wanted %d but got %d", want, got)
}

// NewArgumentCountError :
func NewArgumentCountError(want, got int) *ArgumentCountError {
	return &ArgumentCountError{want, got}
}
