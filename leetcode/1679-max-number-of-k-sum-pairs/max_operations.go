package leetcode

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	answer := 0
	for left < right {
		s := nums[left] + nums[right]
		if s > k {
			right--
		} else if s < k {
			left++
		} else {
			left++
			right--
			answer++
		}
	}
	return answer
}
