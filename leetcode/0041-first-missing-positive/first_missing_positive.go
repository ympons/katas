package leetcode

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for 0 < nums[i] && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func firstMissingPositiveOld(nums []int) int {
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
