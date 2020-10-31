package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n, sum, visited := len(nums), 0, false
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			s := nums[i] + nums[left] + nums[right]
			if !visited || abs(target-s) < abs(target-sum) {
				sum = s
				visited = true
			}
			if s > target {
				right--
			} else if s < target {
				left++
			} else {
				return s
			}
		}
	}
	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
