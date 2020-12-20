package leetcode

func sortedSquares(nums []int) []int {
	n := len(nums)
	l, r, i := 0, n-1, n-1
	squares := make([]int, n)
	for l <= r {
		if left, right := abs(nums[l]), abs(nums[r]); left > right {
			squares[i] = left * left
			i--
			l++
		} else {
			squares[i] = right * right
			i--
			r--
		}
	}
	return squares
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
