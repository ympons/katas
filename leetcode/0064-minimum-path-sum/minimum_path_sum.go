package leetcode

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	D := make([][]int, m)
	for i := 0; i < m; i++ {
		D[i] = make([]int, n)
	}

	D[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		D[i][0] = D[i-1][0] + grid[i][0]
	}

	for j := 1; j < n; j++ {
		D[0][j] = D[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			D[i][j] = grid[i][j]
			if D[i-1][j] < D[i][j-1] {
				D[i][j] += D[i-1][j]
			} else {
				D[i][j] += D[i][j-1]
			}
		}
	}
	return D[m-1][n-1]
}
