package main // import "taylz.io/log/cmd/taylz.io.log.test"

import (
	"fmt"

	"taylz.io/log"
	"taylz.io/log/cmd/taylz.io-log-test/x"
	"taylz.io/log/cmd/taylz.io-log-test/x/examplelongpackagename"
	"taylz.io/log/cmd/taylz.io-log-test/xyz"
)

func main() {
	fmt.Println("-- Color Test --")
	logger := log.StdOutService(log.LevelTrace, log.DefaultFormatWithColor())
	logger.New().Trace("example trace")
	logger.New().Debug("example debug")
	logger.New().Info("example info")
	logger.New().Warn("example warning")
	logger.New().Error("example error")
	logger.New().Out("example out")

	fmt.Println("-- Non-Color Test --")
	logger = log.StdOutService(log.LevelTrace, log.DefaultFormatWithoutColor())
	logger.New().Trace("example trace")
	logger.New().Debug("example debug")
	logger.New().Info("example info")
	logger.New().Warn("example warning")
	logger.New().Error("example error")
	logger.New().Out("example out")

	logger = log.StdOutService(log.LevelTrace, log.DefaultFormatWithColor())
	fmt.Println(`-- CutSourcePackage Test --`)
	logger.New().Debug()
	examplelongpackagename.Info(logger.New())
	x.ClosureInfo(logger.New())
	xyz.Hi(logger.New())
	x.CutSourcePath(logger)
	fmt.Println(`/x.CutSourcePath(logger)`)
	logger.New().Debug()
	examplelongpackagename.Info(logger.New())
	x.ClosureInfo(logger.New())
	xyz.Hi(logger.New())
}
