package log

import "os"

// DailyRollingService creates a Service that writes a rotating log file, named by the day
func DailyRollingService(level Level, f *Format, path string) *Service {
	return NewService(level, f, NewRoller(path))
}

// StdOutService creates a log service that wraps std out
func StdOutService(level Level, f *Format) *Service {
	return NewService(level, f, os.Stdout)
}

// QuickStdout creates a quick-n-dirty log service that wraps std out
func QuickStdout(level Level, name string) *Service {
	return NewService(level, QuickFormat(name), os.Stdout)
}
