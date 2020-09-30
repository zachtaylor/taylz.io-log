package x

import "taylz.io/log"

func ClosureInfo(e log.Logger) {
	func() {
		e.Add("Hello", "World").Info("message content")
	}()
}
