package x

import (
	"fmt"

	"taylz.io/log/logfmt"
	"taylz.io/types"
)

func CutSourcePath(src *logfmt.SourceCutter) {
	fmt.Println(types.NewSource(0), "cut!")
	src.Cut()
}
