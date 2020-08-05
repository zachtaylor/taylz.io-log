package log

// NewEntry creates a new Entry
func NewEntry(flush EntryFunc, fields Fields) *Entry {
	if fields == nil {
		fields = make(Fields)
	}
	return &Entry{
		flush:  flush,
		fields: fields,
	}
}

// Entry is an intermediate step in creating a log
type Entry struct {
	flush  EntryFunc
	fields Fields
}

// EntryFunc is a hook to flush the entry
type EntryFunc = func(Level, Fields, []interface{})

// Add writes any value to the Fields
func (log *Entry) Add(k string, v interface{}) *Entry {
	log.fields[k] = v
	return log
}

// With writes any value to the Fields
func (log *Entry) With(fields Fields) *Entry {
	for k, v := range fields {
		log.fields[k] = v
	}
	return log
}

// Copy duplicates the Entry
func (log *Entry) Copy() *Entry {
	fields := make(Fields)
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
