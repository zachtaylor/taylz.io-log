package logfmt

import (
	"fmt"
	"strconv"
)

// RightPadMinimum returns a formatter that grows strings to minimum size, with right padding
func RightPadMinimum(size int) func(string) string {
	fmts := "%-" + strconv.FormatInt(int64(size), 10) + "s"
	return func(str string) string {
		return fmt.Sprintf(fmts, str)
	}
}
