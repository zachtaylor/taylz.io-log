package x

import "taylz.io/log"

func ClosureInfo(e *log.Entry) {
	func() {
		e.Add("Hello", "World").Info("message content")
	}()
}
