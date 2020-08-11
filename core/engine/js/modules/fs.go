package modules

import (
	"bufio"
	"io"
	"os"
	"path/filepath"

	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto"
)

// FsModule :
type FsModule struct {
	baseDir string
	vm      *otto.Otto
}

// NewFsModule :
func NewFsModule() *FsModule {
	return &FsModule{"", otto.New()}
}

func (f *FsModule) readFile(call otto.FunctionCall) otto.Value {
	filename := call.Argument(0).String()
	handler := call.Argument(1)
	if !handler.IsFunction() {
		panic(f.vm.MakeTypeError("Expected function but found" + handler.Class()))
	}
	filename = filepath.Join(f.baseDir, filename)
	file, err := os.Open(filename)
	if err != nil {
		panic(f.vm.MakeCustomError("Can not read file", err.Error()))
	}
	readFileLineByLine(file, func(line string) {
		lineVal, _ := otto.ToValue(line)
		handler.Call(handler, lineVal)
	})
	return otto.UndefinedValue()
}

// Export :
func (f *FsModule) Export() *otto.Object {
	f.baseDir = lib.FileDirPath
	o := lib.GetEmptyObject(f.vm)
	o.Set("readFile", f.readFile)
	return o
}

func readFileLineByLine(file io.Reader, handler func(string)) {
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		handler(string(line))
	}
}
