package drivers

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/robertkrimen/otto"
)

func TestEventDriver(t *testing.T) {
	flag := false
	t.Run("AddEvent", func(t *testing.T) {
		c := qt.New(t)
		vm := otto.New()
		handler, err := vm.ToValue(func(a interface{}) {
			flag = true
		})
		c.Assert(err, qt.IsNil)
		EventDriver.AddEvent("test", handler)
		c.Assert(len(EventDriver.events), qt.Equals, 1)
	})

	t.Run("Emit", func(t *testing.T) {
		c := qt.New(t)
		EventDriver.Emit("test", otto.UndefinedValue())
		c.Assert(flag, qt.IsTrue)
	})
}
