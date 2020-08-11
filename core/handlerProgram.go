package core

import (
	"strings"

	"github.com/ahmetcanozcan/eago/common/constants"
	"github.com/ahmetcanozcan/eago/common/eagrors"
	"github.com/ahmetcanozcan/eago/core/lib"
	"github.com/robertkrimen/otto/ast"
)

var (
	errDeclaredHandler = eagrors.NewError("Handler program already declared")
)

type handlerBundle struct {
	URLPath        *lib.URLPath
	defaultProgram *ast.Program
	programs       map[string]*ast.Program
}

func newHandlerBundle(urlPath *lib.URLPath) *handlerBundle {
	return &handlerBundle{urlPath, nil, make(map[string]*ast.Program)}
}

func (h *handlerBundle) hasDefault() bool {
	return h.defaultProgram != nil
}

func (h *handlerBundle) addProgram(method string, program *ast.Program) error {
	method = strings.ToUpper(method)
	if method == constants.HTTP.ALL {
		if h.defaultProgram != nil {
			return errDeclaredHandler
		}
		h.defaultProgram = program
		return nil
	}
	if _, ok := h.programs[method]; ok {
		return errDeclaredHandler
	}
	h.programs[method] = program
	return nil
}

func (h *handlerBundle) getProgram(method string) (*ast.Program, error) {
	program, ok := h.programs[method]
	if ok {
		return program, nil
	}
	if h.defaultProgram != nil {
		return h.defaultProgram, nil
	}
	return nil, eagrors.NewError("Handler program not found")
}

// parseHandlers read all script in given dir
// and compile and evalute it to handlerBundler array
func parseHandlers(dir string) (map[string]*handlerBundle, error) {

	scripts, err := parseScriptsFromDir(dir)
	if err != nil {
		return nil, eagrors.NewErrorWithCause(err, "can not parse handlers")
	}
	bundles, err := evaluateSciptsToBundles(scripts)
	if err != nil {
		return nil, eagrors.NewErrorWithCause(err, "can not parse handlers")
	}

	handlerBundles, err := evaluateBundlesToHandlerBundles(bundles)

	if err != nil {
		return nil, eagrors.NewErrorWithCause(err, "can not parse handlers")
	}

	return handlerBundles, nil
}

func evaluateBundlesToHandlerBundles(bundles map[string]*lib.Bundle) (map[string]*handlerBundle, error) {
	mapper := make(map[string]*handlerBundle, 0)
	for path, bundle := range bundles {
		url, method, err := lib.ParseURLFromFilename(path)
		if err != nil {
			return nil, eagrors.NewErrorWithCause(err, "can not parse "+path)
		}
		urlPath := lib.NewURLPath(url)
		if _, ok := mapper[url]; !ok {
			mapper[url] = newHandlerBundle(urlPath)
		}
		err = mapper[url].addProgram(method, bundle.Program)
		if err != nil {
			return nil, eagrors.NewErrorWithCause(err, "can not parse "+path)
		}
	}
	return mapper, nil
}
