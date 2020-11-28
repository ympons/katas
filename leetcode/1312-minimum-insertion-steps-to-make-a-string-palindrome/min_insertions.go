package leetcode

import "strings"

func minInsertions(s string) int {
	n := len(s)
	var builder strings.Builder
	for i := n - 1; i >= 0; i-- {
		builder.WriteByte(s[i])
	}
	r := builder.String()
	return n - lcs(s, r)
}

func lcs(s, r string) int {
	n := len(s)
	D := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		D[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == r[j-1] {
				D[i][j] = D[i-1][j-1] + 1
			} else {
				D[i][j] = D[i-1][j]
				if D[i][j-1] > D[i][j] {
					D[i][j] = D[i][j-1]
				}
			}
		}
	}
	return D[n][n]
}
