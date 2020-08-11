package engine

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/ahmetcanozcan/eago/common/loggers"
	"github.com/ahmetcanozcan/eago/core/lib"
	qt "github.com/frankban/quicktest"
	"github.com/valyala/fasthttp"
)

func TestBaseHandlerRuntime(t *testing.T) {
	loggers.InitializeLoggers()
	t.Run("Request Object", func(t *testing.T) {
		c := qt.New(t)
		port := 1234
		defer startServerOnPort(t, port, func(ctx *fasthttp.RequestCtx) {
			vm := GetHandlerRuntime(ctx, HandlerRuntimeInfo{
				Params: map[string]string{"test": "param"},
			})
			_, err := vm.Run(`
				var m = request.method;
				var p = request.params["test"]; 
				var h = request.header("test");
			`)

			c.Assert(err, qt.Equals, err)

			method := lib.ToStringFromVM(vm, "m", "fail")
			param := lib.ToStringFromVM(vm, "p", "")
			header := lib.ToStringFromVM(vm, "h", "")

			c.Assert(header, qt.Equals, "header")
			c.Assert(method, qt.Not(qt.Equals), "fail")
			c.Assert(param, qt.Equals, "param")

		}).Close()
		url := fmt.Sprintf("http://localhost:%d", port)
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("test", "header")
		_, err := client.Do(req)
		c.Assert(err, qt.IsNil)
	})

	t.Run("Respone Object", func(t *testing.T) {
		c := qt.New(t)
		port := 1235
		defer startServerOnPort(t, port, func(ctx *fasthttp.RequestCtx) {
			vm := GetHandlerRuntime(ctx, HandlerRuntimeInfo{})

			_, err := vm.Run(`
				response.status(300)
				response.write("Test Message")
			`)

			c.Assert(err, qt.Equals, err)

		}).Close()
		url := fmt.Sprintf("http://localhost:%d", port)
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		c.Assert(err, qt.IsNil)
		body, _ := ioutil.ReadAll(resp.Body)
		code := resp.StatusCode
		c.Assert(string(body), qt.Equals, "Test Message")
		c.Assert(code, qt.Equals, code)
	})

}

func startServerOnPort(t *testing.T, port int, h fasthttp.RequestHandler) io.Closer {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		t.Fatalf("cannot start tcp server on port %d: %s", port, err)
	}
	go fasthttp.Serve(ln, h)
	return ln
}
