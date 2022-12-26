package logange

import (
	"fmt"
	"os"
	"time"
)

// StderrHandler is an implementation of the Handler interface that directs to stderr
type StderrHandler struct {
	formatter Formatter
	logLevel  Level
}

// SetFormatter sets the formatter for the handler to use
func (h *StderrHandler) SetFormatter(f Formatter) {
	h.formatter = f
}

// SetLevel sets the level the handler uses
func (h *StderrHandler) SetLevel(lvl Level) {
	h.logLevel = lvl
}

// Level returns the current logging level
func (h *StderrHandler) Level() Level {
	return h.logLevel
}

// LevelString returns the current logging level as a string
func (h *StderrHandler) LevelString() string {
	return LevelToString(h.logLevel)
}

// RecordLog records the log to the stderr stream
func (h *StderrHandler) RecordLog(message string, i []interface{}, logLvl Level, lineno string, name string, datetime time.Time) {
	if logLvl >= h.logLevel {
		fmt.Fprint(os.Stderr, h.formatter.Format(fmt.Sprintf(message, i...), LevelToString(logLvl), lineno, name, datetime))
	}
}

// NewStderrHandler returns an instantiated StderrHandler type
func NewStderrHandler() StderrHandler {
	s := StderrHandler{}
	s.SetFormatter(defaultFormatter)
	return s
}
