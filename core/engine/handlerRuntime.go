package engine

import (
	"net/url"

	"github.com/ahmetcanozcan/eago/common/eagrors"
	"github.com/ahmetcanozcan/eago/common/loggers"
	"github.com/ahmetcanozcan/eago/core/engine/js/bootstrap"
	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto"
	"github.com/valyala/fasthttp"
)

var baseHandlerRuntime *otto.Otto

// GetHandlerRuntime creates a copy of base runtime
// and add response and request object on it
// then return it
func GetHandlerRuntime(ctx *fasthttp.RequestCtx, opt HandlerRuntimeInfo) *otto.Otto {
	vm := baseHandlerRuntime.Copy()
	err := eagrors.HandleErrors(
		vm.Set("request", getRequestObject(ctx, vm, opt)),
		vm.Set("response", getResponseObject(ctx, vm, opt)),
	)
	if err != nil {
		loggers.Debug().Error("GetHandlerRuntime", err)
		// TODO:
	}
	vm.Run(bootstrap.HandlerBootstrapProgram)
	return vm
}

func getRequestObject(ctx *fasthttp.RequestCtx, vm *otto.Otto, opt HandlerRuntimeInfo) otto.Value {
	r := requestObject{handlerObject{vm, opt, ctx}}
	obj := lib.GetEmptyObject(vm)
	obj.Set("method", string(ctx.Method()))
	obj.Set("params", opt.Params)
	obj.Set("header", r.headerFunc)
	obj.Set("body", r.getBody())
	obj.Set("query", r.getQueryParams())
	return obj.Value()
}

type handlerObject struct {
	vm  *otto.Otto
	opt HandlerRuntimeInfo
	ctx *fasthttp.RequestCtx
}

type requestObject struct {
	handlerObject
}

func (r *requestObject) headerFunc(name string) string {
	return string(r.ctx.Request.Header.Peek(name))
}

func (r *requestObject) getBody() *otto.Object {
	o := lib.GetEmptyObject(r.vm)
	o.Set("text", func(call otto.FunctionCall) otto.Value {
		v, _ := otto.ToValue(string(r.ctx.Request.Body()))
		return v
	})
	return o
}

func (r *requestObject) getQueryParams() otto.Value {
	queryStr := string(r.ctx.URI().QueryString())
	queryStr, _ = url.QueryUnescape(queryStr)
	o, _ := r.vm.ToValue(lib.GetQueryValues(queryStr))
	return o
}

func getResponseObject(ctx *fasthttp.RequestCtx, vm *otto.Otto, opt HandlerRuntimeInfo) otto.Value {
	r := responseObject{handlerObject{vm, opt, ctx}, false}
	obj := lib.GetEmptyObject(vm)
	obj.Set("write", r.write)
	obj.Set("status", r.status)
	obj.Set("end", r.end)
	obj.Set("setHeader", r.setHeader)
	obj.Set("__redirect", r.redirect)
	return obj.Value()
}

type responseObject struct {
	handlerObject
	closed bool
}

func (r *responseObject) write(data string) int {
	if !r.closed {
		i, _ := r.ctx.WriteString(data)
		return i
	}
	return -1
}
func (r *responseObject) end() {
	r.closed = true
}

func (r *responseObject) status(code int) {
	r.ctx.SetStatusCode(code)
}

func (r *responseObject) redirect(url string, code int) {
	r.ctx.Redirect(url, code)
}

func (r *responseObject) setHeader(key, value string) {
	r.ctx.Response.Header.Set(key, value)
}

// HandlerRuntimeInfo contains information for handler runtime
type HandlerRuntimeInfo struct {
	// URL paramaters. By default it is nil
	Params map[string]string
}

func (h *HandlerRuntimeInfo) fillDefaults() {
	if h.Params == nil {
		h.Params = make(map[string]string)
	}
}

func init() {
	baseHandlerRuntime = createBaseRuntime()
}
