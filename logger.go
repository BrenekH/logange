package logange

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// Logger is the basic unit of logging
type Logger struct {
	name     string
	handlers []Handler
	parents  []*Logger
}

// AddHandler adds the provided handler to the Logger
func (l *Logger) AddHandler(h Handler) {
	l.handlers = append(l.handlers, h)
}

// AddParent adds a parent logger
func (l *Logger) AddParent(lo *Logger) {
	l.parents = append(l.parents, lo)
}

// NewLogger creates a new logger instance and sets this logger as a parent.
func (l *Logger) NewLogger(name string) Logger {
	newLogger := getLogger(name)
	newLogger.AddParent(l)
	return newLogger
}

// log logs with a specified level
func (l *Logger) log(lvl Level, s string, i []interface{}) {
	_, _, lineno, ok := runtime.Caller(2)
	if !ok {
		lineno = -1
	}

	now := time.Now()
	linenoStr := fmt.Sprint(lineno)

	// Call handlers
	for _, v := range l.handlers {
		v.RecordLog(s, i, lvl, linenoStr, l.name, now)
	}

	// Pass up lineage chain
	for _, v := range l.parents {
		v.parentLog(s, i, lvl, linenoStr, l.name, now)
	}
}

// parentLog logs with all of the parameters passed to it
func (l *Logger) parentLog(message string, i []interface{}, logLvl Level, lineno string, name string, datetime time.Time) {
	// Call handlers
	for _, v := range l.handlers {
		v.RecordLog(message, i, logLvl, lineno, name, datetime)
	}

	// Pass up lineage chain
	for _, v := range l.parents {
		v.parentLog(message, i, logLvl, lineno, name, datetime)
	}
}

// Trace records a log record using a format string with level of Trace.
func (l *Logger) Trace(s string, i ...interface{}) {
	l.log(LevelTrace, s, i)
}

// Debug records a log record using a format string with level of Debug
func (l *Logger) Debug(s string, i ...interface{}) {
	l.log(LevelDebug, s, i)
}

// Info records a log record using a format string with level of Info
func (l *Logger) Info(s string, i ...interface{}) {
	l.log(LevelInfo, s, i)
}

// Warn records a log record using a format string with level of Warn
func (l *Logger) Warn(s string, i ...interface{}) {
	l.log(LevelWarn, s, i)
}

// Error records a log record using a format string with level of Error
func (l *Logger) Error(s string, i ...interface{}) {
	l.log(LevelError, s, i)
}

// Critical records a log record using a format string with level of Critical and then calls os.Exit(1).
func (l *Logger) Critical(s string, i ...interface{}) {
	l.log(LevelCritical, s, i)
	os.Exit(1)
}

// NewLogger returns an properly instantiated Logger type
func NewLogger(name string) Logger {
	l := getLogger(name)

	l.AddParent(RootLogger)

	return l
}

// getLogger returns an properly instantiated Logger type
func getLogger(name string) Logger {
	return Logger{
		name:     name,
		handlers: make([]Handler, 0),
		parents:  make([]*Logger, 0),
	}
}
