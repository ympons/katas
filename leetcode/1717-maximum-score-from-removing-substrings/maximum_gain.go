package leetcode

func maximumGain(s string, x int, y int) int {
	sx, sy := "ab", "ba"
	if x < y {
		sx, sy = sy, sx
		x, y = y, x
	}
	b, gainx := remove([]byte(s), sx, x)
	_, gainy := remove(b, sy, y)
	return gainx + gainy
}

func remove(s []byte, r string, x int) ([]byte, int) {
	gain, i := 0, 0
	for j := 0; j < len(s); j++ {
		s[i] = s[j]
		if i > 0 && s[i-1] == r[0] && s[i] == r[1] {
			gain += x
			i--
		} else {
			i++
		}
	}
	return s[:i], gain
}
