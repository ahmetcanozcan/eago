package modules

import (
	"github.com/ahmetcanozcan/eago/core/drivers"
	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto"
)

// SharedData :
type SharedData struct {
	vm     *otto.Otto
	driver *drivers.SharedData
}

// NewSharedData :
func NewSharedData() *SharedData {
	return &SharedData{otto.New(), drivers.NewSharedData()}
}

// Export :
func (s *SharedData) Export() *otto.Object {
	obj := lib.GetEmptyObject(s.vm)
	obj.Set("set", func(call otto.FunctionCall) otto.Value {
		name := call.Argument(0).String()
		val := call.Argument(1)
		s.driver.Set(name, val)
		return otto.UndefinedValue()
	})
	obj.Set("get", func(call otto.FunctionCall) otto.Value {
		name := call.Argument(0).String()
		return s.driver.Get(name)
	})
	return obj
}
