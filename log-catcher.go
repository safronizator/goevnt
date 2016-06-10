package goevnt

import (
	"bufio"
	"io"
	"log"
)

// LogCatcher cathces log messages and emit event on each message
// TODO: example
type LogCatcher struct {
	BaseEmitter
	namespace string
}

// Creates LogCatcher instance
func NewLogCatcher() *LogCatcher {
	return &LogCatcher{
		BaseEmitter: NewEmitter(),
	}
}

// Creates LogCatcher instance fith namespace prefix
func NewNamespacedLogCatcher(namespace string) *LogCatcher {
	return &LogCatcher{
		BaseEmitter: NewNamespacedEmitter(namespace),
	}
}

// Use logger, created by this method, to catch messages from
func (c *LogCatcher) MakeLogger(eventName, prefix string, flag int) *log.Logger {
	reader, writer := io.Pipe()
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			c.Emit(eventName, scanner.Text())
		}
	}()
	return log.New(writer, prefix, flag)
}

// Creates logger with std log flags. 'log' used as event name by default
func (c *LogCatcher) MakeStdLogger() *log.Logger {
	return c.MakeLogger("log", "", log.LstdFlags)
}
