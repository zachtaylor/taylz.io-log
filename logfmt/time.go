package logfmt

import "taylz.io/types"

// TimeDefault uses 24-hour time format ("15:04:05")
func TimeDefault(time types.Time) string { return time.Format("15:04:05") }

// TimeWipe formats a time to clear content
func TimeWipe(time types.Time) string { return "" }

// TimeReplace returns a formatter that reuses the time slot of the log string to print something else
func TimeReplace(str string) func(types.Time) string { return func(t types.Time) string { return str } }
