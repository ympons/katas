package leetcode

func maxSubArray(nums []int) int {
	max, current := -(1 << 31), 0
	for _, x := range nums {
		current += x
		if current < x {
			current = x
		}
		if max < current {
			max = current
		}
	}
	return max
}
