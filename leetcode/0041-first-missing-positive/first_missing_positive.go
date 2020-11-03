package leetcode

func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	n := len(nums)
	for i, v := range nums {
		if v < 0 || v >= n {
			nums[i] = 0
		}
	}
	for _, v := range nums {
		nums[v%n] += n
	}
	for i, v := range nums {
		if v/n == 0 {
			return i
		}
	}
	return n
}
