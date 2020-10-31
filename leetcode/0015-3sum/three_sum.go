package leetcode

import "sort"

// O(n^2)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	triplets := make([][]int, 0, n/3)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		target := -nums[i]
		for left < right {
			if s := nums[left] + nums[right]; s > target {
				right--
			} else if s < target {
				left++
			} else {
				triplets = append(triplets, []int{nums[i], nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}

	}
	return triplets
}
