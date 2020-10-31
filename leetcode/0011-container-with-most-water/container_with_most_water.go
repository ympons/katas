package leetcode

// two pointer technique - O(n)
func maxArea(height []int) int {
	area, l, r := 0, 0, len(height)-1
	for l < r {
		area = max(area, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return area
}

// naive/brute force solution - O(n^2)
func maxAreaNaive(height []int) int {
	n, area := len(height), 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			area = max(area, (j-i)*min(height[i], height[j]))
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
