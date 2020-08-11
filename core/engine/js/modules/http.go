package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto"
)

var methodMap = map[string]string{
	"get":    "GET",
	"post":   "POST",
	"put":    "PUT",
	"delete": "DELETE",
}

// HTTP :
type HTTP struct {
	vm *otto.Otto
}

// NewHTTP :
func NewHTTP() *HTTP {
	return &HTTP{otto.New()}
}

func (h *HTTP) requestFunc(call otto.FunctionCall) otto.Value {
	method := call.Argument(0).String()
	url := call.Argument(1).String()
	body := ""
	if !call.Argument(2).IsUndefined() {
		body = call.Argument(2).String()
	}
	resp, err := makeHTTPRequest(method, url, body)
	if err != nil {
		return otto.UndefinedValue()
	}
	return h.getResponseObject(resp).Value()
}

func (h *HTTP) getResponseObject(resp *http.Response) *otto.Object {
	o := lib.GetEmptyObject(h.vm)
	o.Set("status", resp.StatusCode)
	nytes, _ := ioutil.ReadAll(resp.Body)
	o.Set("body", string(nytes))
	o.Set("header", func(call otto.FunctionCall) otto.Value {
		name := call.Argument(0).String()
		header := resp.Header.Get(name)
		return lib.ToValueFromString(header)
	})
	return o
}

// Export :
func (h *HTTP) Export() *otto.Object {
	o := lib.GetEmptyObject(h.vm)
	o.Set("request", h.requestFunc)
	o.Set("methods", methodMap)
	return o
}

func makeHTTPRequest(method, url, body string) (*http.Response, error) {
	client := &http.Client{}
	bodyReader := strings.NewReader(body)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		fmt.Println("err", err)
	}
	return client.Do(req)
}
