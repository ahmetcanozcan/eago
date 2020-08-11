package modules

import (
	"strconv"
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
	res := lib.GetEmptyObject(s.vm)
	for i, row := range rows {
		rowObj, err := s.vm.ToValue(row)
		if err != nil {
			panic(err)
		}
		ind := strconv.Itoa(i)
		res.Set(ind, rowObj)
	}
	return res.Value()
}

func isSelectQuery(query string) bool {
	query = strings.TrimLeft(query, "\t\n\r ")
	head := query[:len("SELECT")]
	return head == "SELECT" || head == "select"
}

// NewSQLModule :
func NewSQLModule() *SQL {
	return &SQL{otto.New(), drivers.NewSQLDatabaseDriver()}
}

// Export :
func (s *SQL) Export() *otto.Object {
	val, _ := s.vm.ToValue(func(call otto.FunctionCall) otto.Value {
		_s := NewSQLModule()
		call.This.Object().Set("exec", _s.execQuery)
		call.This.Object().Set("connect", _s.getConnectFunc())
		call.This.Object().Set("disconnect", _s.disconnect)
		return call.This
	})
	return val.Object()
}
