# goevnt
Package goevnt implements publish-subscribe pattern very simple manner.
Just embed BaseEmitter in your structs to let them emit events.

## Disclaimer
Lib is under active development. Some API may be added or changed.

## Example:
```go
import (
  "fmt"
  "time"
  "github.com/safronizator/goevnt"
)

type SomeCounter struct {
	goevnt.BaseEmitter
	i int
}

func NewSomeCounter() *SomeCounter {
	return &SomeCounter{
		BaseEmitter: goevnt.NewEmitter(),
	}
}

func (c *SomeCounter) Start() {
	for {
		time.Sleep(3 * time.Second)
		c.i++
		c.Emit("tick", c.i)
	}
}

func main() {
	c := NewSomeCounter()
	c.Subscribe("tick", goevnt.NewHandlerFunc(func(e goevnt.Event) {
		fmt.Println("Counter tick:", e.Data())
	}))
	c.Start()
}
```

## TODOs:
- Subscribe to event group (with wildcards)
- Tests
- More examples
