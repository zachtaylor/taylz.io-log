package log

import "taylz.io/types"

// NewEntry creates a new Entry
func NewEntry(flush EntryFunc, fields types.Dict) *Entry {
	if fields == nil {
		fields = make(types.Dict)
	}
	return &Entry{
		flush:  flush,
		fields: fields,
	}
}

// Entry is an intermediate step in creating a log
type Entry struct {
	flush  EntryFunc
	fields types.Dict
}

// EntryFunc is a hook to flush the entry
type EntryFunc = func(Level, types.Dict, []interface{})

// Add writes any value to the types.Dict
func (log *Entry) Add(k string, v interface{}) *Entry {
	log.fields[k] = v
	return log
}

// With writes any value to the types.Dict
func (log *Entry) With(fields types.Dict) *Entry {
	for k, v := range fields {
		log.fields[k] = v
	}
	return log
}

// Copy duplicates the Entry
func (log *Entry) Copy() *Entry {
	fields := make(types.Dict)
	for k, v := range log.fields {
		fields[k] = v
	}
	return &Entry{
		flush:  log.flush,
		fields: fields,
	}
}

// Trace attempts to flush a log with LevelTrace
func (log *Entry) Trace(args ...interface{}) { log.flush(LevelTrace, log.fields, args) }

// Debug attempts to flush a log with LevelDebug
func (log *Entry) Debug(args ...interface{}) { log.flush(LevelDebug, log.fields, args) }

// Info attempts to flush a log with LevelInfo
func (log *Entry) Info(args ...interface{}) { log.flush(LevelInfo, log.fields, args) }

// Warn attempts to flush a log with LevelWarn
func (log *Entry) Warn(args ...interface{}) { log.flush(LevelWarn, log.fields, args) }

// Error attempts to flush a log with LevelError
func (log *Entry) Error(args ...interface{}) { log.flush(LevelError, log.fields, args) }

// Out attempts to flush a log with LevelOut
func (log *Entry) Out(args ...interface{}) { log.flush(LevelOut, log.fields, args) }
