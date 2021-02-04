package logange

import (
	"time"
)

// RootLogger is the parent of all Loggers
var RootLogger *Logger

func init() {
	rL := getLogger("root")
	RootLogger = &rL
}

// Handler is an interface for the Handlers used by logging
type Handler interface {
	RecordLog(message string, logLvl Level, lineno string, name string, datetime time.Time)
	SetLevel(level Level)
	SetFormatter(f Formatter)
}

// Formatter is an interface for the Formatter used by logging
type Formatter interface {
	Format(message string, levelname string, lineno string, name string, datetime time.Time) string
}
