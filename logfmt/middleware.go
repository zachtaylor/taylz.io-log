package logfmt

// Middleware returns a formatter that takes 2 formatters, and runs them in sequence
func Middleware(this, next func(string) string) func(string) string {
	return func(str string) string {
		if str != "" {
			str = this(str)
		}
		if str != "" {
			str = next(str)
		}
		return str
	}
}
