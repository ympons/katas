package leetcode

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	X := make([][]int, m)
	for i := 0; i < m; i++ {
		X[i] = make([]int, n)
	}
	A := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a := 0
			if i-1 >= 0 {
				a = X[i-1][j]
			}
			b := 0
			if j-1 >= 0 {
				b = X[i][j-1]
			}
			c := 0
			if i-1 >= 0 && j-1 >= 0 {
				c = X[i-1][j-1]
			}
			X[i][j] = matrix[i][j] ^ a ^ b ^ c
			A = append(A, X[i][j])
		}
	}
	sort.Ints(A)
	return A[len(A)-k]
}
