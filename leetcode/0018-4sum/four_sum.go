package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	quadruplets, n := [][]int{}, len(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				s := nums[i] + nums[j] + nums[left] + nums[right]
				if s > target {
					right--
				} else if s < target {
					left++
				} else {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
				}
			}
		}
	}
	return quadruplets
}
