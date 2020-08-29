package drivers

import (
	"errors"

	"github.com/robertkrimen/otto"
)

var (
	// EventDriver :
	EventDriver = &eventDriver{make(map[string]*event)}
)

type eventDriver struct {
	events map[string]*event
}

type event struct {
	handler otto.Value
	evalCh  chan struct {
		payload otto.Value
		cb      chan otto.Value
	}
}

// Emit sends given event and returns resposne of that event
func (ed *eventDriver) Emit(name string, payload otto.Value) (otto.Value, error) {
	ev, ok := ed.events[name]
	if !ok {
		return otto.UndefinedValue(), errors.New("Event " + name + " not found")
	}
	cb := make(chan otto.Value)
	ev.evalCh <- struct {
		payload otto.Value
		cb      chan otto.Value
	}{
		payload,
		cb,
	}
	return <-cb, nil
}

// AddEvent :
func (ed *eventDriver) AddEvent(name string, handler otto.Value) {
	ed.events[name] = &event{
		handler: handler,
		evalCh: make(chan struct {
			payload otto.Value

			cb chan otto.Value
		}),
	}
	go ed.events[name].run()
}

func (e *event) run() {
	go func() {
		for {
			pv := <-e.evalCh
			res, err := e.handler.Call(e.handler, pv.payload)
			if err != nil {
				panic(err)
			}
			pv.cb <- res
		}
	}()
}
