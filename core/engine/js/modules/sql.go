package modules

import (
	"strings"

	"github.com/ahmetcanozcan/eago/core/drivers"
	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto"
)

// SQL native sql module for db connection and crud operations
type SQL struct {
	vm     *otto.Otto
	driver *drivers.SQLDatabaseDriver
}

func (s *SQL) getConnectFunc() interface{} {
	return func(dsn string) bool {
		err := s.driver.Connect("postgres", dsn)
		if err != nil {
			err := s.vm.MakeCustomError("Connection Error", err.Error())
			panic(err)
		}
		return true
	}
}

func (s *SQL) disconnect(call otto.FunctionCall) otto.Value {
	s.driver.Disconnect()
	return otto.TrueValue()
}

func (s *SQL) execQuery(call otto.FunctionCall) otto.Value {
	query := call.Argument(0).String()
	if len(query) < len("SELECT ") {
		panic(s.vm.MakeCustomError("Invalid Query", query))
	}
	var val otto.Value

	if isSelectQuery(query) {
		res, err := s.driver.ExecuteSelectQuery(query)
		if err != nil {
			panic(s.vm.MakeTypeError(err.Error()))
		}
		val = s.getOttoValueFromRowSlice(res)

	} else {
		res, err := s.driver.ExecuteQuery(query)
		if err != nil {
			panic(s.vm.MakeTypeError(err.Error()))
		}
		val, _ = otto.ToValue(res)
	}
	return val
}

func (s *SQL) getOttoValueFromRowSlice(rows []drivers.SQLRow) otto.Value {
	v, _ := s.vm.ToValue(rows)
	return v
}

func isSelectQuery(query string) bool {
	query = strings.TrimLeft(query, "\t\n\r ")
	head := query[:len("SELECT")]
	return strings.ToUpper(head) == "SELECT"
}

// NewSQLModule :
func NewSQLModule() *SQL {
	return &SQL{otto.New(), drivers.NewSQLDatabaseDriver()}
}

// Export :
func (s *SQL) Export() *otto.Object {
	o := lib.GetEmptyObject(s.vm)
	o.Set("Postgres", func(call otto.FunctionCall) otto.Value {
		_s := NewSQLModule()
		dsn := call.Argument(0)
		if dsn.IsString() {
			if err := _s.driver.Connect("postgres", dsn.String()); err != nil {
				panic(_s.vm.MakeCustomError("DB CONNECTION ERROR", err.Error()))
			}
		}
		call.This.Object().Set("exec", _s.execQuery)
		call.This.Object().Set("connect", _s.getConnectFunc())
		call.This.Object().Set("disconnect", _s.disconnect)
		return call.This
	})
	return o
}
