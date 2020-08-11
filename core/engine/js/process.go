package js

import (
	"os"
	"strings"

	"github.com/ahmetcanozcan/eago/common/eagrors"
	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto"
)

// Process is the global object that includes utilities about process
type Process struct {
}

// NewProcess :
func NewProcess() *Process {
	return &Process{}
}

// Exit exits Javascript process with exiting code
func (p *Process) Exit(call otto.FunctionCall) otto.Value {
	exitCode := lib.ToIntFromValue(call.Argument(0), 0)
	// This panic is only recovered on function that  vm started
	panic(eagrors.NewProcessExitError(exitCode))
}

// Export :
func (p *Process) Export(vm *otto.Otto) {
	obj := lib.GetEmptyObject(vm)
	obj.Set("env", getEnvObject(vm))
	obj.Set("exit", p.Exit)
	vm.Set("process", obj)
}

// getEnvObjec : returns a json object that contains environment variables
func getEnvObject(vm *otto.Otto) *otto.Object {
	o := lib.GetEmptyObject(vm)
	for _, env := range os.Environ() {
		ar := strings.Split(env, "=")
		key, value := ar[0], ar[1]
		o.Set(key, lib.ToValueFromString(value))
	}
	return o
}
