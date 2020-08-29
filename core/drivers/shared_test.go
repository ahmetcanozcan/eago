package drivers

import (
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
	"github.com/robertkrimen/otto"
)

func TestSharedDta(t *testing.T) {
	sd := NewSharedData()
	t.Run("Get", func(t *testing.T) {
		c := qt.New(t)
		sd.values["test1"] = newSharedObject(otto.TrueValue())
		val := sd.Get("test1")
		b, _ := val.ToBoolean()
		c.Assert(b, qt.IsTrue)
	})

	t.Run("Set", func(t *testing.T) {
		c := qt.New(t)
		sd.Set("test2", otto.TrueValue())
		time.Sleep(200 * time.Millisecond)
		val, _ := sd.values["test2"].val.ToBoolean()
		c.Assert(val, qt.IsTrue)
	})

	t.Run("Update", func(t *testing.T) {
		c := qt.New(t)
		val, _ := otto.ToValue(1)
		sd.values["updateTest"] = newSharedObject(val)
		sd.Update("updateTest", func(v otto.Value) otto.Value {
			num, _ := v.ToInteger()
			res, _ := otto.ToValue(num + 1)
			return res
		})
		n, _ := sd.Get("updateTest").ToInteger()
		c.Assert(n, qt.Equals, int64(2))
	})
}
