package drivers

import "github.com/robertkrimen/otto"

// SharedData :
type SharedData struct {
	values   map[string]otto.Value
	updateCh chan struct {
		name  string
		value otto.Value
	}
	getCh chan struct {
		name string
		cb   chan otto.Value
	}
}

// NewSharedData :
func NewSharedData() *SharedData {
	s := &SharedData{
		make(map[string]otto.Value),
		make(chan struct {
			name  string
			value otto.Value
		}),
		make(chan struct {
			name string
			cb   chan otto.Value
		}),
	}
	s.run()
	return s
}

// Get :
func (sd *SharedData) Get(name string) otto.Value {
	cb := make(chan otto.Value)
	sd.getCh <- struct {
		name string
		cb   chan otto.Value
	}{name, cb}
	return <-cb
}

// Set sets a shared value
func (sd *SharedData) Set(name string, val otto.Value) {
	sd.updateCh <- struct {
		name  string
		value otto.Value
	}{name, val}
}

func (sd *SharedData) run() {
	go func() {
		for {
			select {
			case pair := <-sd.updateCh:
				sd.values[pair.name] = pair.value
			case pair := <-sd.getCh:
				pair.cb <- sd.values[pair.name]
			}
		}
	}()
}
