package drivers

import (
	"sync"

	"github.com/robertkrimen/otto"
)

type sharedObject struct {
	val      otto.Value
	getMutex *sync.Mutex
	setMutex *sync.Mutex
}

func newSharedObject(v otto.Value) *sharedObject {
	return &sharedObject{
		val:      v,
		getMutex: &sync.Mutex{},
		setMutex: &sync.Mutex{},
	}
}

// SharedData :
type SharedData struct {
	values map[string]*sharedObject
	setCh  chan struct {
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
		make(map[string]*sharedObject),
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
	sd.setCh <- struct {
		name  string
		value otto.Value
	}{name, val}
}

// Update get and sets the data in same operation
func (sd *SharedData) Update(name string, cb func(otto.Value) otto.Value) {
	_, ok := sd.values[name]
	if ok {

		sd.values[name].getMutex.Lock()
		sd.values[name].setMutex.Lock()

		sd.values[name].val = cb(sd.values[name].val)

		sd.values[name].getMutex.Unlock()
		sd.values[name].setMutex.Unlock()

	} else {
		val := cb(otto.UndefinedValue())
		sd.Set(name, val)
	}
}

func (sd *SharedData) run() {

	go func() {
		for {
			select {
			case pair := <-sd.setCh:
				_, ok := sd.values[pair.name]
				if ok {
					sd.values[pair.name].setMutex.Lock()
					sd.values[pair.name].val = pair.value
					sd.values[pair.name].setMutex.Unlock()
				} else {
					sd.values[pair.name] = newSharedObject(pair.value)
				}
			case pair := <-sd.getCh:
				obj, ok := sd.values[pair.name]
				if ok {
					obj.getMutex.Lock()
					pair.cb <- obj.val
					obj.getMutex.Unlock()
				} else {
					pair.cb <- otto.UndefinedValue()
				}
			}
		}
	}()
}
