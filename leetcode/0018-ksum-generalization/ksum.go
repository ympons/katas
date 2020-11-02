package leetcode

import "sort"

// O(n^(k-1))
func kSum(nums []int, target, k int) [][]int {
	sort.Ints(nums)
	return kSumAuxiliar(nums, target, k)
}

func kSumAuxiliar(nums []int, target, k int) [][]int {
	output, n := [][]int{}, len(nums)
	if n == 0 || nums[0]*k > target || nums[n-1]*4 < target {
		return output
	}
	if k == 2 {
		return twoSum(nums, target)
	}
	for i := 0; i < n; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			for _, set := range kSumAuxiliar(nums[i+1:], target-nums[i], k-1) {
				output = append(output, append([]int{nums[i]}, set...))
			}
		}
	}
	return output
}

func twoSum(nums []int, target int) [][]int {
	output, n := [][]int{}, len(nums)
	left, right := 0, n-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target || (left > 0 && nums[left] == nums[left-1]) {
			left++
		} else if sum > target || (right < n-1 && nums[right] == nums[right+1]) {
			right--
		} else {
			output = append(output, []int{nums[left], nums[right]})
			left++
			right--
		}
	}
	return output
}
