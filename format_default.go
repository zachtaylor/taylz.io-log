package log

import (
	"fmt"
	"time"
)

// DefaultFormatWithColor creates a new Format with colors, and default time and message formats
func DefaultFormatWithColor() *Format {
	f := NewFormat(DefaultTimeFormat, DefaultSourceFormat, DefaultMessageFormat, DefaultColors())
	f.CutPathWith(NewSource(1), 0)
	return f
}

// DefaultFormatWithoutColor creates a new Format without colors, and default time and message formats
func DefaultFormatWithoutColor() *Format {
	f := NewFormat(DefaultTimeFormat, DefaultSourceFormat, DefaultMessageFormat, nil)
	f.CutPathWith(NewSource(1), 0)
	return f
}

// DefaultTimeFormat returns "15:04:05" (24-hour format)
func DefaultTimeFormat(time time.Time) string {
	return time.Format("15:04:05")
}

// DefaultSourceFormat formats a string to max length 32, elipses the beginning
func DefaultSourceFormat(src string) string {
	const maxlen = 32
	lensrc := len(src)
	if lensrc <= maxlen {
		return src
	}
	lendif := lensrc - maxlen
	buf := make([]byte, maxlen)
	buf[0] = '.'
	buf[1] = '.'
	buf[2] = '.'
	for i := 3; i < maxlen; i++ {
		buf[i] = src[lendif+i]
	}
	return string(buf)
}

// DefaultMessageFormat formats a string to "%-15s " (minimum length 16, left-padded, and last char is space)
func DefaultMessageFormat(msg string) string {
	return fmt.Sprintf("%-15s ", msg)
}
