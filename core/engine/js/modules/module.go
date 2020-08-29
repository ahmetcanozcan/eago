package modules

import (
	"github.com/robertkrimen/otto"
)

// systemModules without prefix
var systemModules = map[string]SysModule{
	"fs":     NewFsModule(),
	"psql":   NewSQLModule(),
	"shared": NewSharedData(),
	"http":   NewHTTP(),
	"events": NewEventModule(),
}

// GetSystemModule returns system module
func GetSystemModule(name string) (SysModule, bool) {
	val, ok := systemModules[name]
	return val, ok
}

// SysModule :
type SysModule interface {
	Export() *otto.Object
}
