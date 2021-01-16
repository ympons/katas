package leetcode

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}
	nums, max := make([]int, n+1), 0
	nums[1] = 1
	for i := 1; i <= n/2; i++ {
		nums[2*i] = nums[i]
		if max < nums[i] {
			max = nums[i]
		}
		if 2*i+1 > n {
			return max
		}
		nums[2*i+1] = nums[i] + nums[i+1]
		if max < nums[2*i+1] {
			max = nums[2*i+1]
		}
	}
	return max
}
