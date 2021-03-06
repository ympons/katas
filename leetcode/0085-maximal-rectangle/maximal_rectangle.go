package leetcode

// monotonic queue
func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	if m == 0 {
		return 0
	}
	H, answer := make([]int, m+1), 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				H[j]++
			} else {
				H[j] = 0
			}
		}
		deque := []int{}
		for j := 0; j <= m; j++ {
			for len(deque) > 0 && H[deque[len(deque)-1]] >= H[j] {
				height := H[deque[len(deque)-1]]
				deque = deque[:len(deque)-1]
				width := j
				if len(deque) > 0 {
					width = j - 1 - deque[len(deque)-1]
				}
				if area := height * width; answer < area {
					answer = area
				}
			}
			deque = append(deque, j)
		}
	}
	return answer
}

func maximalRectangleDP(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	if m == 0 {
		return 0
	}

	L, R, H := make([]int, m), make([]int, m), make([]int, m)
	for j := 0; j < m; j++ {
		R[j] = m - 1
	}

	maxArea := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				H[j]++
			} else {
				H[j] = 0
			}
		}

		left := 0
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				if L[j] < left {
					L[j] = left
				}
			} else {
				left, L[j] = j+1, 0
			}
		}

		right := m - 1
		for j := m - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if R[j] > right {
					R[j] = right
				}
			} else {
				right, R[j] = j-1, m-1
			}
		}

		for j := 0; j < m; j++ {
			if area := (R[j] - L[j] + 1) * H[j]; maxArea < area {
				maxArea = area
			}
		}
	}

	return maxArea
}
