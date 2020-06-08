package x

import "taylz.io/log"

func CutSourcePath(logger *log.Service) *log.Service {
	logger.Format().CutPathSource()
	return logger
}
