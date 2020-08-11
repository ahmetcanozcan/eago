package engine

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestBaseRuntime(t *testing.T) {
	t.Run("Error Handling", func(t *testing.T) {
		c := qt.New(t)
		baseRuntime = createBaseRuntime()
		vm := GetBaseRuntime()
		_, err := vm.Run(`
		var v = "test"	
		try {
				var c =	__require("nonExistFile");
				v = "non-test"
			} catch(er) {			}
		`)
		strVal, _ := vm.Get("v")
		str := strVal.String()
		c.Assert(str, qt.Equals, "test")
		c.Assert(err, qt.IsNil)
	})
}
