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

// DefaultSourceFormat formats a left-padded string with length 24, elipses the beginning if necessary
func DefaultSourceFormat(src string) string {
	const finlen = 24
	lensrc := len(src)
	lendif := lensrc - finlen
	buf := make([]byte, finlen)
	i := 0
	j := 0
	if lendif > 0 {
		buf[i] = '.'
		i++
		buf[i] = '.'
		i++
		buf[i] = '.'
		i++
		j = lendif + 3
	}
	for i < finlen && j < lensrc {
		buf[i] = src[j]
		i++
		j++
	}
	for i < finlen {
		buf[i] = ' '
		i++
	}
	return string(buf)
}

// DefaultMessageFormat formats a string to "%-18s" (minimum length 18, right-padded)
func DefaultMessageFormat(msg string) string {
	return fmt.Sprintf("%-18s", msg)
}
