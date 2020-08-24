package examplelongpackagename

import "taylz.io/log"

func Info(e *log.Entry) {
	e.Add("Hello", "World").Info("example longer message content")
}
