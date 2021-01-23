package leetcode

import "sort"

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	sortDiagonal := func(row, col int) {
		D := []int{}
		for row < m && col < n {
			D = append(D, mat[row][col])
			row++
			col++
		}
		sort.Ints(D)
		for k := len(D) - 1; k >= 0; k-- {
			row--
			col--
			mat[row][col] = D[k]
		}
	}

	for i := 0; i < n; i++ {
		sortDiagonal(0, i)
	}

	for j := 1; j < m; j++ {
		sortDiagonal(j, 0)
	}

	return mat
}
