package logange

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// FileHandler is an implementation of the Handler interface that directs to stdout
type FileHandler struct {
	formatter Formatter
	logLevel  Level
	file      *os.File
}

// SetFormatter sets the formatter for the handler to use
func (h *FileHandler) SetFormatter(f Formatter) {
	h.formatter = f
}

// SetLevel sets the level the handler uses
func (h *FileHandler) SetLevel(lvl Level) {
	h.logLevel = lvl
}

// Level returns the current logging level
func (h *FileHandler) Level() Level {
	return h.logLevel
}

// LevelString returns the current logging level as a string
func (h *FileHandler) LevelString() string {
	return LevelToString(h.logLevel)
}

// RecordLog records the log to the stdout stream
func (h *FileHandler) RecordLog(message string, i []interface{}, logLvl Level, lineno string, name string, datetime time.Time) {
	if logLvl >= h.logLevel {
		_, err := h.file.Write([]byte(h.formatter.Format(fmt.Sprintf(message, i...), LevelToString(logLvl), lineno, name, datetime)))
		if err != nil {
			fmt.Printf("Error writing to log file %v: %v\n", h.file.Name(), err)
		}
	}
}

// Close closes the underlying os.File object
func (h *FileHandler) Close() {
	h.file.Close()
}

// NewFileHandler returns an instantiated FileHandler type
func NewFileHandler(path string) (FileHandler, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0644); err != nil {
		return FileHandler{}, err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return FileHandler{}, err
	}

	return FileHandler{formatter: defaultFormatter, logLevel: LevelInfo, file: f}, nil
}
