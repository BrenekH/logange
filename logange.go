package logange

// RootLogger is the parent of all Loggers
var RootLogger *Logger

// defaultFormatter is a basic formatter applied by default
var defaultFormatter Formatter

func init() {
	rL := getLogger("root")
	RootLogger = &rL

	defaultFormatter = StandardFormatter{FormatString: "${datetime}|${name}|${levelname} ${message}\n"}
}
