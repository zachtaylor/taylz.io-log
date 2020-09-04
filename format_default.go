package log

import (
	"taylz.io/log/logfmt"
	"taylz.io/types"
)

// NewSourceCuttingFormatter creates a `taylz.io/log/logfmt/SourceCutter`
func NewSourceCuttingFormatter() *logfmt.SourceCutter { return logfmt.NewSourceCutter() }

// DefaultSourceFormatter creates a formatter with the direct call stack parent package added to the source cut set
func DefaultSourceFormatter() func(string) string {
	return NewSourceCuttingFormatter().CutPathWith(types.NewSource(1), 0).Format
}

// DefaultFormatWithColor creates a new Format with colors, and default time and message formats
func DefaultFormatWithColor() *Format {
	return NewFormat(DefaultTimeFormatter, logfmt.Middleware(
		NewSourceCuttingFormatter().CutPathWith(types.NewSource(1), 0).Format,
		logfmt.RightPadLeftElide(20),
	), DefaultMessageFormatter, DefaultColors())
}

// DefaultFormatWithoutColor creates a new Format without colors, and default time and message formats
func DefaultFormatWithoutColor() *Format {
	return NewFormat(DefaultTimeFormatter, logfmt.Middleware(
		NewSourceCuttingFormatter().CutPathWith(types.NewSource(1), 0).Format,
		logfmt.RightPadLeftElide(20),
	), DefaultMessageFormatter, nil)
}

// QuickFormat creates a Format for 1-off executables / scripts that replaces the time field with the given prefix, and default time and message formats
func QuickFormat(str string) *Format {
	return NewFormat(ReplaceTimeFormatter(str), WipeFormatter, DefaultMessageFormatter, DefaultColors())
}

// ClampFormatter returns a formatter that produces string of set size, right-padded or elided if necessary
var ClampFormatter = logfmt.RightPadLeftElide

// DefaultTimeFormatter uses 24-hour time format ("15:04:05")
var DefaultTimeFormatter = logfmt.TimeDefault

// DefaultMessageFormatter formats a string to "%-20s" (minimum length 20, right-padded)
var DefaultMessageFormatter = logfmt.RightPadMinimum(20)

// ReplaceTimeFormatter returns a formatter that reuses the time slot of the log string to print something else
var ReplaceTimeFormatter = logfmt.TimeReplace

// MiddlewareFormatter returns a formatter that takes 2 formatters, and runs them in sequence
var MiddlewareFormatter = logfmt.Middleware

// WipeTimeFormatter formats a time to clear content
var WipeTimeFormatter = logfmt.TimeWipe

// WipeFormatter formats a string to clear content
var WipeFormatter = logfmt.Wipe
