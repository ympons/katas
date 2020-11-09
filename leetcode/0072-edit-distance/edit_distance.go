package leetcode

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	D := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		D[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		D[i][0] = i
	}

	for j := 0; j <= n; j++ {
		D[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				D[i][j] = D[i-1][j-1]
			} else {
				D[i][j] = min(D[i-1][j-1], min(D[i][j-1], D[i-1][j])) + 1
			}
		}
	}

	return D[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
