package modules

import (
	"github.com/ahmetcanozcan/eago/core/drivers"
	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto"
)

// EventModule :
type EventModule struct {
	vm *otto.Otto
}

// NewEventModule :
func NewEventModule() *EventModule {
	return &EventModule{otto.New()}
}

// Export :
func (e *EventModule) Export() *otto.Object {
	o := lib.GetEmptyObject(e.vm)

	o.Set("emit", func(call otto.FunctionCall) otto.Value {
		name := call.Argument(0).String()
		payload := call.Argument(1)
		resp, err := drivers.EventDriver.Emit(name, payload)
		if err != nil {
			panic(e.vm.MakeCustomError("Event Errors", err.Error()))
		}
		return resp
	})

	return o
}
