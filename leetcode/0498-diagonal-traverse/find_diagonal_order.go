package leetcode

func findDiagonalOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n, r, c, up := len(matrix[0]), 0, 0, true
	order := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		order[i] = matrix[r][c]
		if up {
			if r == 0 || c == n-1 {
				up = false
				if c == n-1 {
					r++
				} else {
					c++
				}
			} else {
				r--
				c++
			}
		} else {
			if r == m-1 || c == 0 {
				up = true
				if r == m-1 {
					c++
				} else {
					r++
				}
			} else {
				r++
				c--
			}
		}
	}
	return order
}
