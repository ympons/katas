package leetcode

// monotonic queue - O(n)
func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]

	deque := []int{0}
	for i := 1; i < n; i++ {
		for len(deque) > 0 && deque[0]+k < i {
			deque = deque[1:]
		}
		dp[i] = dp[deque[0]] + nums[i]
		for len(deque) > 0 && dp[deque[len(deque)-1]] <= dp[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	return dp[n-1]
}
