package drivers

import (
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
	"github.com/robertkrimen/otto"
)

func TestSharedDta(t *testing.T) {
	sd := NewSharedData()
	sd.values["test1"] = otto.TrueValue()
	t.Run("Get", func(t *testing.T) {
		c := qt.New(t)
		val := sd.Get("test1")
		b, _ := val.ToBoolean()
		c.Assert(b, qt.IsTrue)
	})

	t.Run("Set", func(t *testing.T) {
		c := qt.New(t)
		sd.Set("test2", otto.TrueValue())
		time.Sleep(200 * time.Millisecond)
		val, _ := sd.values["test2"].ToBoolean()
		c.Assert(val, qt.IsTrue)
	})
}
