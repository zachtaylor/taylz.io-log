package x

import "taylz.io/log"

func Hi(e *log.Entry) {
	func() {
		e.Info()
	}()
}
