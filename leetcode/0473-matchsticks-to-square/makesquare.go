package leetcode

import "sort"

func makesquare(nums []int) bool {
	n, sum, max := len(nums), 0, 0
	for _, v := range nums {
		sum += v
		if max < v {
			max = v
		}
	}
	if n < 4 || sum%4 > 0 || max > sum/4 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
	target, groups := sum/4, make([]int, 4)

	var dfs func(int) bool
	dfs = func(p int) bool {
		if p == n {
			return true
		}
		M := map[int]struct{}{}
		for i := 0; i < 4; i++ {
			if _, seen := M[groups[i]]; seen || groups[i]+nums[p] > target {
				continue
			}
			M[groups[i]] = struct{}{}
			groups[i] += nums[p]
			if dfs(p + 1) {
				return true
			}
			groups[i] -= nums[p]
		}
		return false
	}

	return dfs(0)
}
