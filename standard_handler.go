package logange

import (
	"fmt"
	"time"
)

// StdoutHandler is an implementation of the Handler interface that directs to stdout
type StdoutHandler struct {
	formatter Formatter
	logLevel  Level
}

// SetFormatter sets the formatter for the handler to use
func (h *StdoutHandler) SetFormatter(f Formatter) {
	h.formatter = f
}

// SetLevel sets the level the handler uses
func (h *StdoutHandler) SetLevel(lvl Level) {
	h.logLevel = lvl
}

// RecordLog records the log to the stdout stream
func (h *StdoutHandler) RecordLog(message string, logLvl Level, lineno string, name string, datetime time.Time) {
	if logLvl >= h.logLevel {
		fmt.Print(h.formatter.Format(message, LevelToString(logLvl), lineno, name, datetime))
	}
}
