package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	var prefix strings.Builder
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	m := len(strs[0])
	for i := 0; i < m; i++ {
		c, matches := strs[0][i], 1
		for j := 1; j < n && i < len(strs[j]); j++ {
			if strs[j][i] == c {
				matches++
			} else {
				break
			}
		}
		if matches != n {
			break
		}
		prefix.WriteByte(c)
	}
	return prefix.String()
}
