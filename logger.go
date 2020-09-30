package log

import "taylz.io/types"

// Logger is a log writer
type Logger interface {
	New() Entry
	Add(string, interface{}) Entry
	With(types.Dict) Entry
	Trace(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Out(...interface{})
}
