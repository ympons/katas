package leetcode

// D(i,j) = min(D(i-1,j), D(i-1,j-1), D(i,j-1)) + 1
func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	D := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		D[i] = make([]int, m+1)
	}

	size := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if matrix[i-1][j-1] == '1' {
				D[i][j] = min(D[i-1][j], min(D[i-1][j-1], D[i][j-1])) + 1
				if D[i][j] > size {
					size = D[i][j]
				}
			}
		}
	}

	return size * size
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
