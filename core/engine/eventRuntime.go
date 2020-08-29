package engine

import (
	"github.com/ahmetcanozcan/eago/core/drivers"
	"github.com/robertkrimen/otto"
)

var baseEventRuntime *otto.Otto

// GetEventRuntime :
func GetEventRuntime(eventName string) *otto.Otto {
	vm := baseEventRuntime.Copy()

	vm.Set("listen", func(call otto.FunctionCall) otto.Value {
		handler := call.Argument(0)

		if !handler.IsFunction() {
			return otto.UndefinedValue()
		}
		drivers.EventDriver.AddEvent(eventName, handler)
		return otto.UndefinedValue()
	})

	return vm
}

func init() {
	baseEventRuntime = createBaseRuntime()
}
