package leetcode

func flipAndInvertImage(A [][]int) [][]int {
	n, m := len(A), len(A[0])
	for i := 0; i < n; i++ {
		// flip and invert the row
		for j := (m+1)/2 - 1; j >= 0; j-- {
			A[i][j], A[i][m-j-1] = A[i][m-j-1]^1, A[i][j]^1
		}
	}
	return A
}
