package lib

import (
	"github.com/ahmetcanozcan/eago/common/eagrors"
	"github.com/robertkrimen/otto"
)

// CheckForArgumentCount checks for argument count.
// if got and want values are not equal, panic
func CheckForArgumentCount(vm *otto.Otto, call otto.FunctionCall, want int) {
	got := len(call.ArgumentList)
	if want != got {
		panic(eagrors.NewArgumentCountError(want, got).JSError(vm))
	}
}
