package modules

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ahmetcanozcan/eago/common/eagofs"
	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto"
)

const (
	defaultBufferSize = 1024
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

// Reader :
func (f *FsModule) Reader(call otto.FunctionCall) otto.Value {
	if !call.Argument(0).IsString() {
		panic(f.vm.MakeRangeError(
			"want string got " + call.Argument(0).Class()))
	}
	filename, _ := call.Argument(0).ToString()
	filename = filepath.Join(lib.FileDirPath, filename)

	if !eagofs.IsExist(filename) {
		panic(f.vm.MakeCustomError("File Error",
			filename+" doesn't exist"))
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(f.vm.MakeCustomError("File Error", err.Error()))
	}
	reader := bufio.NewReader(file)

	buffer := make([]byte, defaultBufferSize)

	call.This.Object().Set("setBufferSize", func(size int) {
		buffer = make([]byte, size)
	})

	call.This.Object().Set("read",
		func(call otto.FunctionCall) otto.Value {
			n, _ := reader.Read(buffer)
			if n == 0 {
				return otto.UndefinedValue()
			}
			res := buffer[:n]
			val, _ := otto.ToValue(string(res))
			return val
		})
	call.This.Object().Set("close",
		func() {
			file.Close()
		})

	return otto.UndefinedValue()
}

// Writer :
func (f *FsModule) Writer(call otto.FunctionCall) otto.Value {
	if !call.Argument(0).IsString() {
		panic(f.vm.MakeRangeError(
			"want string got " + call.Argument(0).Class()))
	}
	filename, err := call.Argument(0).ToString()
	filename = filepath.Join(lib.FileDirPath, filename)
	if eagofs.IsDir(filename) {
		panic(f.vm.MakeCustomError("Invalid Argument",
			filename+" is not a file"))
	}
	mode := os.O_TRUNC | os.O_WRONLY
	if call.Argument(1).IsString() {
		b, _ := call.Argument(1).ToString()
		if strings.Contains(b, "a") {
			mode = os.O_WRONLY | os.O_APPEND
		}
	}
	var file *os.File
	if eagofs.IsExist(filename) {
		file, err = os.OpenFile(filename, mode, 0644)
		if err != nil {
			panic(f.vm.MakeCustomError("File Error", err.Error()))
		}
	} else {
		file, err = os.Create(filename)
		if err != nil {
			panic(f.vm.MakeCustomError("File Error", err.Error()))
		}
	}
	//writer := bufio.NewWriter(file)

	call.This.Object().Set("write", func(data string) {
		n, err := file.WriteString(data)
		fmt.Println(n)
		if err != nil {
			panic(f.vm.MakeCustomError("File Error",
				err.Error()))
		}
		if n == 0 && data != "" {
			panic(f.vm.MakeCustomError("File Error", "Can not write anything"))
		}
	})

	call.This.Object().Set("close", func() {
		file.Close()
		if err != nil {
			panic(f.vm.MakeCustomError("Can not close Writer", err.Error()))
		}

	})
	return otto.UndefinedValue()
}

func (f *FsModule) readDir(call otto.FunctionCall) otto.Value {
	dirname := call.Argument(0).String()
	if !eagofs.IsDir(dirname) {
		panic(f.vm.MakeCustomError("File Error",
			dirname+" is not a directory"))
	}
	dirname = filepath.Join(f.baseDir, dirname)
	filenames := make([]string, 0)
	eagofs.NewDirReader(dirname).Read(func(filename string) error {
		filenames = append(filenames, filename)
		return nil
	})
	val, _ := f.vm.ToValue(filenames)
	return val
}

// Export :
func (f *FsModule) Export() *otto.Object {
	f.baseDir = lib.FileDirPath
	o := lib.GetEmptyObject(f.vm)

	o.Set("Writer", f.Writer)
	o.Set("Reader", f.Reader)

	o.Set("stat", f.stat)
	o.Set("readDir", f.readDir)
	o.Set("readFile", f.readFile)
	o.Set("writeFile", f.writeFile)
	o.Set("pipe", fsPipeValue(f.vm))
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
