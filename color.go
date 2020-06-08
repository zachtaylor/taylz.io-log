package log

const (
	nocolor = "\x1b[0m"
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	blue    = "\x1b[34m"
	purple  = "\x1b[35m"
	cyan    = "\x1b[36m"
	white   = "\x1b[37m"
)

// DefaultColors returns the default color set
func DefaultColors() map[Level]string {
	return map[Level]string{
		LevelTrace: purple,
		LevelDebug: green,
		LevelInfo:  blue,
		LevelWarn:  yellow,
		LevelError: red,
		LevelOut:   white,
	}
}
