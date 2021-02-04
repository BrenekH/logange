package logange

import (
	"strings"
	"time"
)

// StandardFormatter is the in-library implementation of the Formatter interface
type StandardFormatter struct {
	FormatString string
}

// Format formats using the passed arguments and the formatter's format string
func (f StandardFormatter) Format(message string, levelname string, lineno string, name string, datetime time.Time) (s string) {
	s = f.FormatString

	// Replace items in format string
	// lineno
	s = strings.ReplaceAll(s, "${lineno}", lineno)

	// datetime
	s = strings.ReplaceAll(s, "${datetime}", datetime.String())

	// levelname
	s = strings.ReplaceAll(s, "${levelname}", levelname)

	// name
	s = strings.ReplaceAll(s, "${name}", name)

	// message
	s = strings.ReplaceAll(s, "${message}", message)

	return
}
