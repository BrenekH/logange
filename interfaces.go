package logange

import (
	"time"
)

// Handler is an interface for the Handlers used by logging
type Handler interface {
	RecordLog(message string, i []interface{}, logLvl Level, lineno string, name string, datetime time.Time)
	SetLevel(level Level)
	SetFormatter(f Formatter)
	Level() Level
	LevelString() string
}

// Formatter is an interface for the Formatter used by logging
type Formatter interface {
	Format(message string, levelname string, lineno string, name string, datetime time.Time) string
}
