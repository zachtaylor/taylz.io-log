package logfmt

// RightPadLeftElide returns a formatter that produces string of set size, right-padded or elided if necessary
func RightPadLeftElide(size int) func(string) string {
	return func(src string) string {
		lensrc := len(src)
		lendif := lensrc - size
		buf := make([]byte, size)
		var i, j int
		if lendif > 0 {
			buf[i] = '.'
			i++
			buf[i] = '.'
			i++
			buf[i] = '.'
			i++
			j = lendif + 3
		}
		for i < size && j < lensrc {
			buf[i] = src[j]
			i++
			j++
		}
		for i < size {
			buf[i] = ' '
			i++
		}
		return string(buf)
	}
}
