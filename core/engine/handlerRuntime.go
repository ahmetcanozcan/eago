package engine

import (
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
	err := eagrors.HandleErrors(
		obj.Set("method", r.method()),
		obj.Set("params", r.params()),
		obj.Set("header", r.headerFunc()),
	)
	if err != nil {
		// TODO:
	}
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

func (r *requestObject) method() string {
	return string(r.ctx.Method())
}

func (r *requestObject) params() map[string]string {
	return r.opt.Params
}
func (r *requestObject) headerFunc() interface{} {
	return func(name string) string {
		return string(r.ctx.Request.Header.Peek(name))
	}
}

func getResponseObject(ctx *fasthttp.RequestCtx, vm *otto.Otto, opt HandlerRuntimeInfo) otto.Value {
	r := responseObject{handlerObject{vm, opt, ctx}}
	obj := lib.GetEmptyObject(vm)
	obj.Set("write", r.writeFunc())
	obj.Set("status", r.statusFunc())
	return obj.Value()
}

type responseObject struct {
	handlerObject
}

func (r *responseObject) writeFunc() interface{} {
	return func(data string) int {
		i, _ := r.ctx.WriteString(data)
		return i
	}
}

func (r *responseObject) statusFunc() interface{} {
	return func(code int) {
		r.ctx.SetStatusCode(code)
	}
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
