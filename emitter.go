package goevnt

import (
	"sync"
)

// Base emitter. It can be embedded in your structs, as so as used separate as abstract emitter.
// Implements both of Emitter and Subscriber interfaces
type BaseEmitter struct {
	namespace    string
	subscribers  *nameTree
	subscrLocker *sync.RWMutex
}

// Creates new plain BaseEmitter
func NewEmitter() BaseEmitter {
	return NewNamespacedEmitter("")
}

// Creates BaseEmitter, that prefixes all events with namespace.
func NewNamespacedEmitter(namespace string) BaseEmitter {
	return BaseEmitter{
		namespace:    namespace,
		subscribers:  &nameTree{},
		subscrLocker: &sync.RWMutex{},
	}
}

func (e *BaseEmitter) Subscribe(eventName string, h Handler) {
	e.subscrLocker.Lock()
	defer e.subscrLocker.Unlock()
	var handlers *[]Handler
	if h := e.subscribers.Get(eventName); h != nil {
		handlers = h.Val().(*[]Handler)
	} else {
		hSlice := make([]Handler, 0, 3)
		handlers = &hSlice
		e.subscribers.Set(eventName, handlers)
	}
	*handlers = append(*handlers, h)
}

func (e *BaseEmitter) Emit(eventName string, data interface{}) {
	fullname := eventName
	if e.namespace != "" {
		fullname = e.namespace + "." + eventName
	}
	e.subscrLocker.RLock()
	defer e.subscrLocker.RUnlock()
	hs := e.subscribers.Get(fullname)
	if hs == nil || hs.Val() == nil {
		return
	}
	handlers := hs.Val().(*[]Handler)
	triggredEvent := &EventWrapper{
		name: fullname,
		data: data,
	}
	for _, handler := range *handlers {
		handler.Handle(triggredEvent)
		if triggredEvent.IsPropagationStopped() {
			break
		}
	}
}

////////////////

////////////////

////////////////

////////////////

////////////////
