package leetcode

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	n, m := len(word1), len(word2)
	i, j, ii, jj := 0, 0, 0, 0
	for i < n && j < m {
		if ii < len(word1[i]) && jj < len(word2[j]) && word1[i][ii] != word2[j][jj] {
			return false
		}
		ii++
		jj++
		if ii >= len(word1[i]) {
			ii = 0
			i++
		}
		if jj >= len(word2[j]) {
			jj = 0
			j++
		}
	}
	return i == n && j == m
}

func arrayStringsAreEqualBuffer(word1 []string, word2 []string) bool {
	var b1, b2 strings.Builder

	for _, w1 := range word1 {
		b1.WriteString(w1)
	}

	for _, w2 := range word2 {
		b2.WriteString(w2)
	}

	if b1.Len() != b2.Len() {
		return false
	}

	return b1.String() == b2.String()
}
