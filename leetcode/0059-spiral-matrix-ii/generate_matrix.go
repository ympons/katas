package leetcode

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	row, col, k := 0, 0, n
	direction := 'R'
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		switch direction {
		case 'R':
			col++
			if col == k-1 {
				direction = 'D'
			}
		case 'D':
			row++
			if row == k-1 {
				direction = 'L'
			}
		case 'L':
			col--
			if col == n-k {
				direction = 'U'
			}
		case 'U':
			row--
			if row == n-k {
				direction = 'R'
				k--
				row++
				col++
			}
		}
	}
	return matrix
}
