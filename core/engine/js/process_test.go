package js

import (
	"os"
	"testing"

	"github.com/ahmetcanozcan/eago/core/lib"
	qt "github.com/frankban/quicktest"
	"github.com/robertkrimen/otto"
)

func TestProcess(t *testing.T) {
	
	t.Run("Environment Object", func(t *testing.T) {
		c := qt.New(t)
		
		c.Assert(os.Setenv("Test","Value"),qt.IsNil)

		vm := otto.New()
		NewProcess().Export(vm)
		vm.Run(`
			var a = process.env.Test;
		`)
		envTest := lib.ToStringFromVM(vm,"a","")
		c.Assert(envTest,qt.Equals,"Value")
		c.Assert(os.Setenv("Test",""),qt.IsNil)

	})
} 