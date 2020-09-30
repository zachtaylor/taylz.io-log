package log

import "taylz.io/types"

// Entry is an intermediate step in creating a log
type Entry struct {
	flush  EntryFunc
	fields types.Dict
}

// EntryFunc is a hook to flush the entry
type EntryFunc = func(Level, types.Dict, []interface{})

// NewEntry creates a new Entry
func NewEntry(flush EntryFunc, fields types.Dict) Entry {
	if fields == nil {
		fields = make(types.Dict)
	}
	return Entry{
		flush:  flush,
		fields: fields,
	}
}

func (log *Entry) isLogger() Logger { return log }

// New copies the Entry
func (log Entry) New() Entry {
	copy := NewEntry(log.flush, nil)
	for k, v := range log.fields {
		copy.fields[k] = v
	}
	return copy
}

// Add writes any value to the types.Dict
func (log Entry) Add(k string, v interface{}) Entry {
	copy := log.New()
	copy.fields[k] = v
	return copy
}

// With writes all values to the types.Dict
func (log Entry) With(fields types.Dict) Entry {
	copy := log.New()
	for k, v := range fields {
		copy.fields[k] = v
	}
	return copy
}

// Trace attempts to flush a log with LevelTrace
func (log Entry) Trace(args ...interface{}) { log.flush(LevelTrace, log.fields, args) }

// Debug attempts to flush a log with LevelDebug
func (log Entry) Debug(args ...interface{}) { log.flush(LevelDebug, log.fields, args) }

// Info attempts to flush a log with LevelInfo
func (log Entry) Info(args ...interface{}) { log.flush(LevelInfo, log.fields, args) }

// Warn attempts to flush a log with LevelWarn
func (log Entry) Warn(args ...interface{}) { log.flush(LevelWarn, log.fields, args) }

// Error attempts to flush a log with LevelError
func (log Entry) Error(args ...interface{}) { log.flush(LevelError, log.fields, args) }

// Out attempts to flush a log with LevelOut
func (log Entry) Out(args ...interface{}) { log.flush(LevelOut, log.fields, args) }
