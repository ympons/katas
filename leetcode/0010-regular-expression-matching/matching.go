package leetcode

import "regexp"

// using the Go regexp package
func isMatch(s string, p string) bool {
	pattern := regexp.MustCompile(p)
	if loc := pattern.FindStringIndex(s); loc != nil && loc[0] == 0 && loc[1] == len(s) {
		return true
	}
	return false
}

// naive recursive version
func isMatchRecursive(s, p string) bool {
	n, m := len(s), len(p)
	if m == 0 {
		return n == 0
	}

	firstMatch := false
	if n > 0 && (s[0] == p[0] || p[0] == '.') {
		firstMatch = true
	}

	if m >= 2 && p[1] == '*' {
		return isMatchRecursive(s, p[2:]) || firstMatch && isMatchRecursive(s[1:], p)
	}

	return firstMatch && isMatchRecursive(s[1:], p[1:])
}

// using DP (dynamic programming)
func isMatchDP(s, p string) bool {
	n, m := len(s), len(p)
	if m == 0 {
		return n == 0
	}

	// create a DP table
	DP := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		DP[i] = make([]bool, m+1)
	}

	DP[n][m] = true

	// fill the table in a bottom-up fashion
	for i := n; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			firstMatch := false
			if i < n && (p[j] == s[i] || p[j] == '.') {
				firstMatch = true
			}
			if j+1 < m && p[j+1] == '*' {
				DP[i][j] = DP[i][j+2] || firstMatch && DP[i+1][j]
			} else {
				DP[i][j] = firstMatch && DP[i+1][j+1]
			}
		}
	}

	return DP[0][0]
}
