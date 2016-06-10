package goevnt

// HandlerFunc helps to use closure func as event handler
type HandlerFunc struct {
	f func(e Event)
}

// Creates Handler form closure
func NewHandlerFunc(f func(e Event)) *HandlerFunc {
	return &HandlerFunc{f}
}

func (h *HandlerFunc) Handle(e Event) {
	h.f(e)
}
