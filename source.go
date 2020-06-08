package log

import (
	"runtime"
	"strconv"
)

// Source links to a go code instruction
type Source struct {
	file string
	line int
}

// NewSource creates a Source reference to the caller number # higher in the stack frame than the invoking call (0 is default)
func NewSource(history int) *Source {
	if history < 0 {
		return nil
	}
	_, file, line, _ := runtime.Caller(1 + history)

	return &Source{
		file: file,
		line: line,
	}
}

// File returns Source.file
func (src *Source) File() string {
	return src.file
}

// Line returns Source.line
func (src *Source) Line() int {
	return src.line
}

func (src *Source) String() string {
	msg := ""
	if src != nil {
		if file := src.File(); file != "" {
			if flen := len(file); flen > 4 && file[flen-3:] == ".go" {
				msg = file[:flen-3]
			} else {
				msg = file
			}
		}
		if lno := src.Line(); lno > 0 {
			msg += "#" + strconv.Itoa(lno)
		}
	}
	return msg
}
