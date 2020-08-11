package engine

import (
	"github.com/ahmetcanozcan/eago/common/eagrors"
	"github.com/ahmetcanozcan/eago/core/engine/js"
	"github.com/ahmetcanozcan/eago/core/engine/js/bootstrap"
	"github.com/robertkrimen/otto"
)

// baseRuntime is the parent runtime of all runtimes.
// includes all basic utility and variables.
var baseRuntime *otto.Otto

// GetBaseRuntime returns base runtime
func GetBaseRuntime() *otto.Otto {
	return baseRuntime.Copy()
}

func createBaseRuntime() *otto.Otto {
	vm := otto.New()

	// Set base runtime utilities
	vm.Set("__require", getRequireFunction(vm))

	// Export console.log utilities
	js.NewConsole().Export(vm)
	js.NewProcess().Export(vm)

	// Execute some js scripts for preparing runtime
	vm.Run(bootstrap.EagoBootstrapProgram)
	return vm
}

func getRequireFunction(vm *otto.Otto) interface{} {
	return func(name string) *otto.Object {
		v, err := js.RequireModule(vm, name)
		if err != nil {
			erStr := eagrors.GetErrorString(err)
			panic(vm.MakeCustomError("Can not import", erStr))
		}
		return v
	}
}

func init() {
	baseRuntime = createBaseRuntime()
}
