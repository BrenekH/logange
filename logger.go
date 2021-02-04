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

// AddHandler adds the provided handler to the Logger
func (l *Logger) AddHandler(h Handler) {
	l.handlers = append(l.handlers, h)
}

// AddParent adds a parent logger
func (l *Logger) AddParent(lo *Logger) {
	l.parents = append(l.parents, lo)
}

// log logs with a specified level
func (l *Logger) log(lvl Level, s string) {
	_, _, lineno, ok := runtime.Caller(2)
	if !ok {
		lineno = -1
	}

	now := time.Now()
	linenoStr := fmt.Sprint(lineno)

	// Call handlers
	for _, v := range l.handlers {
		v.RecordLog(s, lvl, linenoStr, l.name, now)
	}

	// Pass up lineage chain
	for _, v := range l.parents {
		v.parentLog(s, lvl, linenoStr, l.name, now)
	}
}

// parentLog logs with all of the parameters passed to it
func (l *Logger) parentLog(message string, logLvl Level, lineno string, name string, datetime time.Time) {
	// Call handlers
	for _, v := range l.handlers {
		v.RecordLog(message, logLvl, lineno, name, datetime)
	}

	// Pass up lineage chain
	for _, v := range l.parents {
		v.parentLog(message, logLvl, lineno, name, datetime)
	}
}

// Trace records a log record with level of Trace
func (l *Logger) Trace(s string) {
	l.log(LevelTrace, s)
}

// Debug records a log record with level of Debug
func (l *Logger) Debug(s string) {
	l.log(LevelDebug, s)
}

// Info records a log record with level of Info
func (l *Logger) Info(s string) {
	l.log(LevelInfo, s)
}

// Warn records a log record with level of Warn
func (l *Logger) Warn(s string) {
	l.log(LevelWarn, s)
}

// Error records a log record with level of Error
func (l *Logger) Error(s string) {
	l.log(LevelError, s)
}

// Critical records a log record with level of Critical and then calls os.Exit(1)
func (l *Logger) Critical(s string) {
	l.log(LevelCritical, s)
	os.Exit(1)
}
