package goevnt

// EventWrapper is a base implementation of Event interface
type EventWrapper struct {
	name               string
	data               interface{}
	propagationStopped bool
}

// DataEvent is key-value container, that helps attach some structured data to events
// TODO: example
type DataEvent struct {
	data map[string]interface{}
}

// Creates new empty DataEvent
func NewDataEvent() *DataEvent {
	return &DataEvent{data: make(map[string]interface{}, 1)}
}

func (e *EventWrapper) StopPropagation() {
	e.propagationStopped = true
}

func (e *EventWrapper) IsPropagationStopped() bool {
	return e.propagationStopped
}

func (e *EventWrapper) Name() string {
	return e.name
}

func (e *EventWrapper) Data() interface{} {
	return e.data
}

func (e *DataEvent) Set(field string, val interface{}) *DataEvent {
	e.data[field] = val
	return e
}

func (e *DataEvent) Get(field string) interface{} {
	data, _ := e.data[field]
	return data
}
