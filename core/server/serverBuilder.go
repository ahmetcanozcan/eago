package server

import (
	"os"
	"path/filepath"
	"time"

	"github.com/ahmetcanozcan/eago/common/eagrors"
	"github.com/ahmetcanozcan/eago/config"
	"github.com/valyala/fasthttp"
)

// LogArguments parameter of  a log function
type LogArguments struct {
	// Execution time of whole request in milliseconds format
	Time   float32
	Status int
	Method string
	Path   string
}

// LogFunc :
type LogFunc func(args *LogArguments)

// Builder helps to building a web-server
type Builder struct {
	handlers []HandlerFunc
	logFunc  LogFunc
}

// NewServerBuilder :
func NewServerBuilder() *Builder {
	return &Builder{}
}

// AddHandlerFunc adds the handler to handler slice
// handlers in the slice will be executed by on by
func (b *Builder) AddHandlerFunc(handler HandlerFunc) *Builder {
	b.handlers = append(b.handlers, handler)
	return b
}

// SetLogFunc sets log function
func (b *Builder) SetLogFunc(f LogFunc) {
	b.logFunc = f
}

// Build server
func (b *Builder) Build() Server {
	s := &fasthttp.Server{
		Name: config.EagoJSON.Name,
	}

	s.Handler = func(ctx *fasthttp.RequestCtx) {
		defer eagrors.RecoverRuntime(string(ctx.Method()), string(ctx.Path()))
		start := time.Now()

		for _, handler := range b.handlers {
			if stop := handler(ctx); stop {
				break
			}
		}

		if b.logFunc != nil {
			elapsed := time.Since(start)
			logArgs := &LogArguments{
				Time:   float32(elapsed.Microseconds()) / 1000.0,
				Method: string(ctx.Method()),
				Path:   string(ctx.Path()),
				Status: ctx.Response.StatusCode(),
			}
			b.logFunc(logArgs)
		}
	}

	return s
}

// FileServerOptions :
type FileServerOptions struct {
	// Root directory of file server
	// Default : "."
	Root string
	// index filenames of fs
	// Default : []string{"index.html"}
	IndexNames []string
}

// NewFileServerHandler :
func NewFileServerHandler(opt FileServerOptions) HandlerFunc {
	opt.filDefaults()
	fs := &fasthttp.FS{
		Root:       opt.Root,
		IndexNames: opt.IndexNames,
	}

	fsHandler := fs.NewRequestHandler()

	return func(ctx *fasthttp.RequestCtx) bool {
		cp := filepath.Join(opt.Root, string(ctx.Path()))
		stat, err := os.Stat(cp)
		if err != nil {
			return false
		}

		if stat.IsDir() {
			_, err := os.Stat(filepath.Join(cp, "index.html"))
			if err != nil {
				return false
			}
			fsHandler(ctx)
			return true
		}
		fsHandler(ctx)
		return true
	}
}

func checkFileExist(filename string, indexNames []string) bool {
	_, err := os.Open(filename)
	if err == nil {
		return true
	}
	for _, index := range indexNames {
		fname := filepath.Join(filename, index)
		_, err := os.Open(fname)
		if err == nil {
			return true
		}
	}

	return false
}

func (opt *FileServerOptions) filDefaults() {
	if opt.IndexNames == nil {
		opt.IndexNames = []string{"index.html"}
	}
	if len(opt.Root) == 0 {
		opt.Root = "."
	}
}
