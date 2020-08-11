package js

import (
	"fmt"
	"testing"

	"github.com/ahmetcanozcan/eago/core/lib"
	qt "github.com/frankban/quicktest"
	"github.com/robertkrimen/otto"
)

func TestWrapAndEvaluateModuleCode(t *testing.T) {
	fmt.Println("HELLO")
	c := qt.New(t)
	code := `
	exports.num = 15;
	exports.name = "testName";
	exports.flag = __dirname == "test/";
	`
	obj := wrapAndEvalModuleCode(otto.New(), code, "test/file.js")

	val, _ := obj.Get("num")
	num, _ := val.ToInteger()

	c.Assert(num, qt.Equals, int64(15))

	val, _ = obj.Get("name")
	name, _ := val.ToString()

	c.Assert(name, qt.Equals, "testName")

	val, _ = obj.Get("flag")
	boolFlag, _ := val.ToBoolean()

	c.Assert(boolFlag, qt.Equals, true)
}

func TestValidatePath(t *testing.T) {
	c := qt.New(t)

	lib.ModuleDirPath = "/test/example"
	testCases := []struct {
		val    string
		result bool
	}{
		{"/test/example/script", true},
		{"/test/not_example/script", false},
	}
	for _, _case := range testCases {
		c.Assert(validateModulePath(_case.val), qt.Equals, _case.result)
	}

}
