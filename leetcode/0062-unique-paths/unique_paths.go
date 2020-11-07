package leetcode

func uniquePaths(m int, n int) int {
	D := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		D[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		D[i][0] = 1
	}
	for j := 0; j <= n; j++ {
		D[0][j] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			D[i][j] = D[i-1][j] + D[i][j-1]
		}
	}

	return D[m-1][n-1]
}
