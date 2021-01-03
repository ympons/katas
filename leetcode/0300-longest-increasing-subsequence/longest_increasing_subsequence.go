package leetcode

import "sort"

func lengthOfLIS(A []int) int {
	D := make([]int, 0, len(A))
	for _, x := range A {
		i := sort.SearchInts(D, x)
		if i == len(D) {
			D = append(D, x)
		} else {
			D[i] = x
		}
	}
	return len(D)
}

func lengthOfLISOld(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	D, answer := make([]int, n), 1
	D[0] = 1
	for i := 1; i < n; i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && D[j] > max {
				max = D[j]
			}
		}
		D[i] = max + 1
		if D[i] > answer {
			answer = D[i]
		}
	}

	return answer
}
