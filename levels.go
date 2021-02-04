package logange

// Level is a type alias for int
type Level int

var (
	// LevelTrace TODO
	LevelTrace Level = 0
	// LevelDebug TODO
	LevelDebug Level = 10
	// LevelInfo TODO
	LevelInfo Level = 20
	// LevelWarn TODO
	LevelWarn Level = 30
	// LevelError TODO
	LevelError Level = 40
	// LevelCritical TODO
	LevelCritical Level = 50
)

var levelMap map[Level]string = map[Level]string{0: "TRACE", 10: "DEBUG", 20: "INFO", 30: "WARNING", 40: "ERROR", 50: "CRITICAL"}

// LevelToString converts the inputted Level id into a string name of the level name
func LevelToString(id Level) (s string) {
	s = ""
	lkp, prs := levelMap[id]
	if prs {
		s = lkp
	}
	return
}
