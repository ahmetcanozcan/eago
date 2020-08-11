// Package eagrors contains common eago erros and related utilities
package eagrors

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"

	_errors "github.com/pkg/errors"
)

// As defined in "github.com/pkg/errors"
type causer interface {
	Cause() error
}

type stackTracer interface {
	StackTrace() _errors.StackTrace
}

// PrintStackTraceFromErr prints the error's stack trace to stdout.
func PrintStackTraceFromErr(err error) {
	FprintStackTraceFromErr(os.Stdout, err)
}

// FprintStackTraceFromErr prints the error's stack trace to w.
func FprintStackTraceFromErr(w io.Writer, err error) {
	if err, ok := err.(stackTracer); ok {
		for _, f := range err.StackTrace() {
			fmt.Fprintf(w, "%+s:%d\n", f, f)
		}
	}
}

// PrintStackTrace prints the current stacktrace to w.
func PrintStackTrace(w io.Writer) {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	fmt.Fprintf(w, "%s", buf)
}

// ErrorSender is a, typically, non-blocking error handler.
type ErrorSender interface {
	SendError(err error)
}

// GetGID gets the current goroutine id. Used only for debugging.
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
