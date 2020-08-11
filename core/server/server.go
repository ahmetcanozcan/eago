package server

import "github.com/valyala/fasthttp"

// Server represent eago core server
type Server interface {
	ListenAndServe(addr string) error
}

// HandlerFunc handles reqeust and return bool
// if it returns true, server stop execution handlers
type HandlerFunc func(*fasthttp.RequestCtx) bool
