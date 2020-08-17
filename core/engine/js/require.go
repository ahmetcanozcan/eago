package js

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ahmetcanozcan/eago/common/eagrors"
	"github.com/ahmetcanozcan/eago/config"
	"github.com/ahmetcanozcan/eago/core/compiler"
	"github.com/ahmetcanozcan/eago/core/engine/js/modules"
	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto"
)

var cachedModules = make(map[string]*otto.Object)

// RequireModule requires user, system or 3rd party  modules
func RequireModule(vm *otto.Otto, name string) (*otto.Object, error) {
	erStr := "Can not require " + name
	// Firstly check for already imported modules
	val, ok := cachedModules[name]
	if ok {
		return val, nil
	}

	// TODO: Set Mutex lock for safe concurrent processing

	mod, ok := modules.GetSystemModule(name)
	if ok {
		val = mod.Export()
		return val, nil
	}
	// Finally check for user module
	var exactPath string
	if filepath.IsAbs(name) {
		exactPath = name
	} else {
		// if module name starts with '@' it refers a 3rd party package
		if strings.HasPrefix(name, "@/") {
			prefix := "@/" + config.EagoJSON.Name
			// if projects is a package then delete
			if config.EagoJSON.Package && strings.HasPrefix(name, prefix) {
				exactPath = filepath.Join(lib.BasePath, name[len(prefix):])
			} else {
				exactPath = filepath.Join(lib.PackageDirPath, name[2:])
			}
		} else {
			exactPath = filepath.Join(lib.ModuleDirPath, name)
		}
	}

	if !strings.HasSuffix(exactPath, ".js") {
		exactPath += ".js"
	}
	_, err := os.Open(exactPath)
	if err != nil {
		// if file is not exist then try for /index.js
		exactPath = filepath.Join(exactPath[:(len(exactPath)-len(".js"))], "index.js")
		_, err = os.Open(exactPath)
		if err != nil {
			return otto.NullValue().Object(), eagrors.NewErrorWithCause(err, erStr)
		}
	}
	// if path it's not valid invoke error
	if !validateModulePath(exactPath) {

	}
	script, err := compiler.ReadScript(exactPath)
	if err != nil {
		return otto.NullValue().Object(), eagrors.NewErrorWithCause(err, erStr)
	}
	obj := wrapAndEvalModuleCode(vm, script.Code, exactPath)
	cachedModules[name] = obj
	return obj, nil
}

func wrapAndEvalModuleCode(vm *otto.Otto, code, exactPath string) *otto.Object {
	_dirname, _filename := filepath.Split(exactPath)
	_dirname = strings.ReplaceAll(_dirname, "\\", "\\\\")

	requireFuncCode := fmt.Sprintf("function require(str) {"+
		"if (str.substr(0, 2) == \"./\") {"+
		"	str = \"%s\" + str.substr(2);}"+
		"return __require(str);	}", _dirname)

	code = fmt.Sprintf(
		`(function(__dirname,__filename,exports,require){ `+
			` %s ; return exports;})("%s","%s",{},%s)`,
		code,
		_dirname,
		_filename,
		requireFuncCode,
	)
	obj, err := vm.Object(code)
	if err != nil {
		panic(err)
	}
	return obj
}

func validateModulePath(path string) bool {
	return strings.HasPrefix(path, lib.ModuleDirPath)
}

type requireLocks struct {
	ls map[string]*sync.Mutex
	l  sync.Mutex
}

func (r *requireLocks) Get(name string) *sync.Mutex {
	_, ok := r.ls[name]
	if !ok {
		r.l.Lock()
		r.ls[name] = &sync.Mutex{}
		r.l.Unlock()
	}
	return r.ls[name]
}

var rlocks = requireLocks{make(map[string]*sync.Mutex), sync.Mutex{}}
