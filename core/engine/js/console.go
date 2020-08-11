package js

import (
	"strings"

	"github.com/ahmetcanozcan/eago/common/loggers"
	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto"
)

type logType int

const (
	infoLogType  = 0
	warnLogType  = 1
	errLogType   = 2
	fatalLogType = 3
)

// Console :
type Console struct {
}

// Info logs info message
func (c *Console) Info(call otto.FunctionCall) otto.Value {
	return log(loggers.Default(), infoLogType, call)
}

// Err logs error message
func (c *Console) Err(call otto.FunctionCall) otto.Value {
	return log(loggers.Default(), errLogType, call)
}

// Warn logs warn message
func (c *Console) Warn(call otto.FunctionCall) otto.Value {
	return log(loggers.Default(), warnLogType, call)
}

// Log logs info message
func (c *Console) Log(call otto.FunctionCall) otto.Value {
	return c.Info(call)
}

// Export exports console as a otto.Value
func (c *Console) Export(vm *otto.Otto) {
	obj := lib.GetEmptyObject(vm)
	obj.Set("log", c.Log)
	obj.Set("error", c.Err)
	obj.Set("warn", c.Warn)
	obj.Set("info", c.Info)
	vm.Set("console", obj)
}

func log(logger loggers.Logger, lType logType, call otto.FunctionCall) otto.Value {
	messages := make([]string, len(call.ArgumentList))
	for i, arg := range call.ArgumentList {
		strVal := arg.String()
		messages[i] = strVal
	}
	message := strings.Join(messages, " ")
	switch lType {
	case errLogType:
		logger.Error(message)
	case warnLogType:
		logger.Warn(message)
	case infoLogType:
		logger.Info(message)
	case fatalLogType:
		logger.Fatal(message)
	}
	return otto.TrueValue()
}

// NewConsole returns new console
func NewConsole() *Console {
	return &Console{}
}
