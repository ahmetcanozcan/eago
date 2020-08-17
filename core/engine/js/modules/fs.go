package modules

import (
	"bufio"
	"io"
	"os"
	"path/filepath"

	"github.com/ahmetcanozcan/eago/common/eagofs"
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

func (f *FsModule) stat(call otto.FunctionCall) otto.Value {
	fname := call.Argument(0).String()
	o := lib.GetEmptyObject(f.vm)
	stat, err := os.Stat(filepath.Join(f.baseDir, fname))
	if err != nil {
		panic(f.vm.MakeCustomError("File stat err", err.Error()))
	}
	o.Set("dir", stat.IsDir())
	o.Set("size", stat.Size())
	return o.Value()
}

func (f *FsModule) writeFile(call otto.FunctionCall) otto.Value {
	filename := call.Argument(0).String()
	data := call.Argument(1).String()
	dirname, filename := filepath.Split(filename)
	dirname = filepath.Join(f.baseDir, dirname)
	absPath := filepath.Join(dirname, filename)
	// Firstly check target dir exist
	dirStat, err := os.Stat(dirname)
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			// if dir is not exist create new one
			eagofs.MkdirAll(dirname, 0777)
		default:
			panic(f.vm.MakeCustomError("File Error", err.Error()))
		}
	}
	// if path is exist but it's not a directory
	if !dirStat.IsDir() {
		panic(f.vm.MakeCustomError("File Error", dirname+" is not directory"))
	}

	file, err := os.Create(absPath)
	if err != nil {
		panic(f.vm.MakeCustomError("File Error", err.Error()))
	}
	_, err = file.WriteString(data)
	if err != nil {
		panic(f.vm.MakeCustomError("File Error", err.Error()))
	}
	return otto.UndefinedValue()
}

// Export :
func (f *FsModule) Export() *otto.Object {
	f.baseDir = lib.FileDirPath
	o := lib.GetEmptyObject(f.vm)
	o.Set("stat", f.stat)
	o.Set("readFile", f.readFile)
	o.Set("writeFile", f.writeFile)
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
