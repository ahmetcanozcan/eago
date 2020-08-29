package modules

import (
	"fmt"
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/robertkrimen/otto"
)

func TestReadFileLineByLine(t *testing.T) {
	c := qt.New(t)
	testStr := []string{
		"This is",
		"multiline text message",
		"for testing",
		"readfile func",
	}

	res := make([]string, len(testStr))
	i := 0
	readFileLineByLine(strings.NewReader(strings.Join(testStr, "\n")), func(line string) {
		res[i] = line
		i++
	})

	c.Assert(res, qt.CmpEquals(), testStr)

}

func TestFSModule(t *testing.T) {
	fs := NewFsModule().Export()
	vm := otto.New()
	vm.Set("fs", fs)
	vm.Set("log", func() interface{} {
		return func(s string) {
			fmt.Println(s)
		}
	}())
	t.Run("pipe", func(t *testing.T) {
		c := qt.New(t)
		tvm := vm.Copy()
		_, err := tvm.Run(`
		var mockReader = (function () {
			var data = "Hello From Test";
			var isSent = false
		
			function read() {
				if(isSent) return;
				isSent= true;
				return data;
			}
			return { read: read };
		})();
		var mockWriter = {};
		var result = "";
		mockWriter.write = function (data) {
			result += data;
		};
		fs.pipe(mockReader,mockWriter)
		`)
		c.Assert(err, qt.IsNil)
		result, err := tvm.Get("result")
		c.Assert(result.String(), qt.Equals, "Hello From Test")
	})

}
