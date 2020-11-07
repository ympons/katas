package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	D := make([][]int, m)
	for i := 0; i < m; i++ {
		D[i] = make([]int, n)
	}

	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		D[i][0] = 1
	}

	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		D[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				D[i][j] = D[i-1][j] + D[i][j-1]
			}
		}
	}

	return D[m-1][n-1]
}
