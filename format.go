package log

import (
	"fmt"

	"taylz.io/types"
)

// Format controls the text output specifics
type Format struct {
	// TimeFmt is the time Format
	TimeFmt func(types.Time) string
	// SourceFmt is the source length
	SourceFmt func(string) string
	// MessageFmt is the message length
	MessageFmt func(string) string
	// Colors is the colors to use per level, or nil when in no-color mode
	Colors map[Level]string
}

// NewFormat creates Format
func NewFormat(tfmt func(types.Time) string, srcfmt, msgfmt func(string) string, colors map[Level]string) *Format {
	return &Format{
		TimeFmt:    tfmt,
		SourceFmt:  srcfmt,
		MessageFmt: msgfmt,
		Colors:     colors,
	}
}

// Format creates writable output
func (f *Format) Format(time types.Time, src *types.Source, lvl Level, flds types.Dict, msg string) []byte {
	var sb types.StringBuilder
	if str := f.TimeFmt(time); len(str) > 0 {
		sb.WriteString(str)
		sb.WriteByte(32) // space
	}
	if f.Colors != nil {
		sb.WriteString(f.Colors[lvl])
	} else {
		sb.WriteByte(lvl.ByteCode())
		sb.WriteByte(32) // space
	}
	if str := f.SourceFmt(src.String()); len(str) > 0 {
		sb.WriteString(str)
		sb.WriteByte(32) // space
	}
	sb.WriteString(f.MessageFmt(msg))
	if f.Colors != nil {
		sb.WriteString(nocolor)
	}
	for _, k := range types.GetKeysDict(flds) {
		fmt.Fprintf(&sb, " %s=%v", k, flds[k])
	}
	sb.WriteByte(10) // newline
	return []byte(sb.String())
}
