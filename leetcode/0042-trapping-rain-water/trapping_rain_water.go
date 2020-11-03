package leetcode

// two pointer technique
func trap(height []int) int {
	n, answer := len(height), 0
	if n == 0 {
		return answer
	}
	left, right, maxLeft, maxRight := 0, n-1, 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				answer += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				answer += maxRight - height[right]
			}
			right--
		}
	}
	return answer
}

// using DP (dynamic programming)
func trapDP(height []int) int {
	n, answer := len(height), 0
	if n == 0 {
		return answer
	}

	maxLeft, maxRight := make([]int, n), make([]int, n)
	maxLeft[0], maxRight[n-1] = height[0], height[n-1]

	for i := 1; i < n; i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
		maxRight[n-i-1] = max(height[n-i-1], maxRight[n-i])
	}

	for i := 0; i < n; i++ {
		answer += min(maxLeft[i], maxRight[i]) - height[i]
	}

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
