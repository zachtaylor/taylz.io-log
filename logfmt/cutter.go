package logfmt

import (
	"path"
	"sort"
	"strings"

	"taylz.io/types"
)

// SourceCutter is a list of strings, with formatter func to be used by log format engine
type SourceCutter struct {
	List []string
}

// NewSourceCutter creates a SourceCutter
func NewSourceCutter() *SourceCutter {
	return &SourceCutter{
		List: make([]string, 0),
	}
}

// Cut adds the caller's package to the Format's PathCut
func (src *SourceCutter) Cut() { src.CutPathWith(types.NewSource(1), 0) }

// CutParent adds the caller's n-recursive parent package to the Format's PathCut
func (src *SourceCutter) CutParent(n int) { src.CutPathWith(types.NewSource(1), n) }

// CutPathWith adds the calling source's path ancestor number to the path cutset
func (src *SourceCutter) CutPathWith(source *types.Source, parentno int) *SourceCutter {
	filePath, _ := path.Split(source.File())
	for i := 0; strings.Contains(filePath, "/") && i <= parentno; i++ {
		filePath = filePath[:strings.LastIndex(filePath, "/")]
	}
	src.List = append(src.List, filePath)
	sort.Slice(src.List, func(i, j int) bool {
		return src.List[i] > src.List[j]
	})
	return src
}

// Format is a func(string) string that cuts source prefixes
func (src *SourceCutter) Format(str string) string {
	for _, path := range src.List {
		if pkglen := len(path); pkglen >= len(str) {
		} else if str[:pkglen] != path {
		} else if str[pkglen] == '/' {
			return str[pkglen+1:]
		}
	}
	return str
}
