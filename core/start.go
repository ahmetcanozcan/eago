package core

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ahmetcanozcan/eago/common/eagrors"
	"github.com/ahmetcanozcan/eago/common/loggers"
	"github.com/ahmetcanozcan/eago/config"
	"github.com/ahmetcanozcan/eago/core/compiler"
	"github.com/ahmetcanozcan/eago/core/engine"
	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/ahmetcanozcan/eago/core/server"
	"github.com/valyala/fasthttp"
)

// Start starts application on given context
func Start(opt StartOptions) {
	opt.fillDefaults()
	config.Parse(opt.AppDir)
	loggers.InitializeLoggers()
	lib.UpdateEnginePathVars(opt.AppDir)

	bundles, err := parseHandlers(lib.HandlerDirPath)

	if err != nil {
		loggers.Default().Error(err)
		return
	}

	if err := runStartJS(filepath.Join(opt.AppDir, "start.js")); err != nil {
		loggers.Default().Error(err)
		return
	}

	builder := server.NewServerBuilder()
	// Handler
	builder.AddHandlerFunc(getHandlerHandlerFunc(bundles))
	if config.EagoJSON.StaticPath != "" {
		// Fileserver
		rootPath := filepath.Join(lib.FileDirPath, config.EagoJSON.StaticPath)
		fs := server.NewFileServerHandler(
			server.FileServerOptions{
				Root: rootPath,
			})
		builder.AddHandlerFunc(fs)
	}
	// 404 handler
	builder.AddHandlerFunc(func(ctx *fasthttp.RequestCtx) bool {
		script, ok := bundles["/404/"]
		if ok {
			engine.GetBaseRuntime().Run(script)
			return false
		}

		file, err := os.Open(filepath.Join(opt.AppDir, "files", config.EagoJSON.NotFound))
		if err == nil {
			ctx.Response.Header.Set("Content-Type", "text/html")
			reader := bufio.NewReader(file)
			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					break
				}
				ctx.Write(line)
			}
			return false
		}
		ctx.SetStatusCode(404)
		ctx.Write([]byte(fmt.Sprintf("Can not %s %s ", string(ctx.Method()), string(ctx.Path()))))
		return false
	})
	builder.SetLogFunc(func(args *server.LogArguments) {
		loggers.Default().Info(
			fmt.Sprintf("%s %s in %.2f ms  %d", args.Method, args.Path, args.Time, args.Status),
		)
	})
	s := builder.Build()
	errCh := make(chan error)
	adr := config.EagoJSON.Address()
	go func() {
		errCh <- s.ListenAndServe(adr)
	}()
	loggers.Default().Info("Listening ", adr, "...")
	err = <-errCh
	if err != nil {
		loggers.Default().Error(err)
	}
}

func getHandlerHandlerFunc(bundles map[string]*handlerBundle) func(ctx *fasthttp.RequestCtx) bool {
	return func(ctx *fasthttp.RequestCtx) bool {
		for _, bundle := range bundles {
			url := string(ctx.Path())
			method := string(ctx.Method())
			if bundle.URLPath.Check(url) {
				prog, err := bundle.getProgram(method)
				if err != nil {
					loggers.Default().Error("Program not found")
				}
				vm := engine.GetHandlerRuntime(ctx, engine.HandlerRuntimeInfo{
					Params: bundle.URLPath.GetURLParams(url),
				})
				_, err = vm.Run(prog)
				if err != nil {
					err = eagrors.NewErrorWithCause(err, fmt.Sprintf("Can not run %s %s", url, method))
					loggers.Default().Error(eagrors.GetErrorString(err))
				}
				return true
			}
		}

		return false
	}
}

func runStartJS(filename string) error {
	script, err := compiler.ReadScript(filename)
	if err != nil {
		return err
	}
	prog, err := compiler.New().Compile(script)
	if err != nil {
		return err
	}
	vm := engine.GetBaseRuntime()
	_, err = vm.Run(prog)
	return err
}
