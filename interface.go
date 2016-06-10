// Package goevnt implements publish-subscribe pattern very simple manner.
// Just embed BaseEmitter in your structs to let them emit events.
//
// Example:
//
// type SomeCounter struct {
// 	goevnt.BaseEmitter
// 	i int
// }

// func NewSomeCounter() *SomeCounter {
// 	return &SomeCounter{
// 		BaseEmitter: goevnt.NewEmitter(),
// 	}
// }

// func (c *SomeCounter) Start() {
// 	for {
// 		time.Sleep(3 * time.Second)
// 		c.i++
// 		c.Emit("tick", c.i)
// 	}
// }

// func main() {
// 	c := NewSomeCounter()
// 	c.Subscribe("tick", goevnt.NewHandlerFunc(func(e goevnt.Event) {
// 		fmt.Println("Counter tick:", e.Data())
// 	}))
// 	c.Start()
// }

package goevnt

// Signals that event propagation was stopped by handler
type PropagationStopper interface {
	StopPropagation()
	IsPropagationStopped() bool
}

// Base event interface. Instances of this are passed to handlers
type Event interface {
	PropagationStopper
	Name() string
	Data() interface{}
}

// Each type that implements this interface can be attached as event handler
type Handler interface {
	Handle(e Event)
}

// Emitter can emit events
type Emitter interface {
	Emit(eventName string, data interface{})
}

// Subscriber can register handlers for events
type Subscriber interface {
	Subscribe(eventName string, h Handler)
}
