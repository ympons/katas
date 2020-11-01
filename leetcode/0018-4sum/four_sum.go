package leetcode

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	quadruplets, n := [][]int{}, len(nums)
	visited := map[string]bool{}
	for i := 0; i < n-3; i++ {
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
					key := fmt.Sprintf("%v|%v|%v|%v", nums[i], nums[j], nums[left], nums[right])
					if !visited[key] {
						quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
						visited[key] = true
					}
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
