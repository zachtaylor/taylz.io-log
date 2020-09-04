package main // import "taylz.io/log/cmd/taylz.io.log.test"

import (
	"fmt"

	"taylz.io/log"
	"taylz.io/log/cmd/taylz.io-log-test/x"
	"taylz.io/log/cmd/taylz.io-log-test/x/examplelongpackagename"
	"taylz.io/log/cmd/taylz.io-log-test/xyz"
)

func main() {
	WithColor()

	NoColor()

	CutSourcePackage()

	QuickFormat()
}

func WithColor() {
	fmt.Println("-- Color Test --")
	logger := log.StdOutService(log.LevelTrace, log.DefaultFormatWithColor())
	logger.New().Trace("example trace")
	logger.New().Debug("example debug")
	logger.New().Info("example info")
	logger.New().Warn("example warning")
	logger.New().Error("example error")
	logger.New().Out("example out")
}

func NoColor() {
	fmt.Println("-- Non-Color Test --")
	logger := log.StdOutService(log.LevelTrace, log.DefaultFormatWithoutColor())
	logger.New().Trace("example trace")
	logger.New().Debug("example debug")
	logger.New().Info("example info")
	logger.New().Warn("example warning")
	logger.New().Error("example error")
	logger.New().Out("example out")
}

func CutSourcePackage() {
	fmt.Println(`-- CutSourcePackage Test --`)
	logsrcfmt := log.NewSourceCuttingFormatter()
	logger := log.StdOutService(log.LevelTrace, log.NewFormat(log.DefaultTimeFormatter, log.MiddlewareFormatter(
		logsrcfmt.Format,
		log.ClampFormatter(32),
	), log.DefaultMessageFormatter, log.DefaultColors()))

	logger.New().Debug()
	examplelongpackagename.Info(logger.New())
	x.ClosureInfo(logger.New())
	xyz.Hi(logger.New())

	x.CutSourcePath(logsrcfmt)

	logger.New().Debug()
	examplelongpackagename.Info(logger.New())
	x.ClosureInfo(logger.New())
	xyz.Hi(logger.New())
}

func QuickFormat() {
	fmt.Println(`-- QuickFormat Test --`)
	logger := log.QuickStdout(log.LevelTrace, "q&d:")

	logger.New().Trace("example trace")
	logger.New().Debug("example debug")
	logger.New().Info("example info")
	logger.New().Warn("example warning")
	logger.New().Error("example error")
	logger.New().Out("example out")
}
