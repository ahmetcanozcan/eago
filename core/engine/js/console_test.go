package js

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/robertkrimen/otto"
)

type dummyLogger struct {
}

func (l dummyLogger) Log(args ...interface{}) {}

func (l dummyLogger) Warn(args ...interface{}) {}

func (l dummyLogger) Error(args ...interface{}) {}

func (l dummyLogger) Fatal(args ...interface{}) {}

func (l dummyLogger) Info(args ...interface{}) {
	first := args[0]
	str, _ := first.(string)
	consoleFlag = str
}

var consoleFlag string = ""

func TestConsole(t *testing.T) {
	consoleFlag = ""
	c := qt.New(t)

	messages := []interface{}{"This", "is", 4, "test"}
	args := make([]otto.Value, len(messages))
	for i, msg := range messages {
		v, _ := otto.ToValue(msg)
		args[i] = v
	}
	call := otto.FunctionCall{
		ArgumentList: args,
	}
	res := log(dummyLogger{}, infoLogType, call)
	b, _ := res.ToBoolean()
	c.Assert(b, qt.Equals, true)
	c.Assert(consoleFlag, qt.Equals, "This is 4 test")
}
