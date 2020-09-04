package log

import (
	"fmt"

	"taylz.io/types"
)

// Service provides logging functionality
type Service struct {
	level Level
	f     *Format
	w     types.WriteCloser
}

// NewService creates a log service with the minimum Level, format function and output dest
func NewService(lvl Level, f *Format, w types.WriteCloser) *Service {
	return &Service{
		level: lvl,
		f:     f,
		w:     w,
	}
}

// New returns a new Entry
func (svc *Service) New() *Entry { return NewEntry(svc.flush, nil) }

// Add returns a new Entry with a field value preset
func (svc *Service) Add(k string, v interface{}) *Entry { return NewEntry(svc.flush, types.Dict{k: v}) }

// With returns a new Entry with the given types.Dict
func (svc *Service) With(fields types.Dict) *Entry { return NewEntry(svc.flush, fields) }

// Trace attempts to flush a log with LevelTrace
func (svc *Service) Trace(args ...interface{}) { svc.flush(LevelTrace, nil, args) }

// Debug attempts to flush a log with LevelDebug
func (svc *Service) Debug(args ...interface{}) { svc.flush(LevelDebug, nil, args) }

// Info attempts to flush a log with LevelInfo
func (svc *Service) Info(args ...interface{}) { svc.flush(LevelInfo, nil, args) }

// Warn attempts to flush a log with LevelWarn
func (svc *Service) Warn(args ...interface{}) { svc.flush(LevelWarn, nil, args) }

// Error attempts to flush a log with LevelError
func (svc *Service) Error(args ...interface{}) { svc.flush(LevelError, nil, args) }

// Out attempts to flush a log with LevelOut
func (svc *Service) Out(args ...interface{}) { svc.flush(LevelOut, nil, args) }

// flush creates a Time and Source to trigger write, only to be used by exposed funcs
func (svc *Service) flush(lvl Level, flds types.Dict, args []interface{}) {
	if lvl >= svc.level {
		svc.w.Write(svc.f.Format(types.NewTime(), types.NewSource(2), lvl, flds, parseargs(args)))
	}
}

// Format returns the format settings
func (svc *Service) Format() *Format { return svc.f }

// Close closes the internal writer
func (svc *Service) Close() error { return svc.w.Close() }

func parseargs(args []interface{}) string {
	var sb = &types.StringBuilder{}
	fmt.Fprint(sb, args...)
	return sb.String()
}
