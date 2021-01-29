package leetcode

import "strings"

func getSmallestString(n int, k int) string {
	var b strings.Builder
	for i := 1; i <= n; i++ {
		size, c := n-i, 1
		if c < k-26*size {
			c = k - 26*size
		}
		k -= c
		b.WriteByte(byte(c + 'a' - 1))
	}
	return b.String()
}
