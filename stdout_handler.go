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

// Level returns the current logging level
func (h *StdoutHandler) Level() Level {
	return h.logLevel
}

// LevelString returns the current logging level as a string
func (h *StdoutHandler) LevelString() string {
	return LevelToString(h.logLevel)
}

// RecordLog records the log to the stdout stream
func (h *StdoutHandler) RecordLog(message string, i []any, logLvl Level, lineno string, name string, datetime time.Time) {
	if logLvl >= h.logLevel {
		fmt.Print(h.formatter.Format(fmt.Sprintf(message, i...), LevelToString(logLvl), lineno, name, datetime))
	}
}

// NewStdoutHandler returns an instantiated StdoutHandler type
func NewStdoutHandler() StdoutHandler {
	s := StdoutHandler{}
	s.SetFormatter(defaultFormatter)
	return s
}
