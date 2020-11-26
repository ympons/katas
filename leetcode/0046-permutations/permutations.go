package leetcode

var output [][]int

func permute(nums []int) [][]int {
	output = [][]int{}
	helper(nums, 0, len(nums))
	return output
}

func helper(nums []int, p int, n int) {
	if p == n {
		tmp := make([]int, n)
		copy(tmp, nums)
		output = append(output, tmp)
		return
	}

	for i := p; i < n; i++ {
		nums[i], nums[p] = nums[p], nums[i]
		helper(nums, p+1, n)
		nums[i], nums[p] = nums[p], nums[i]
	}
}
