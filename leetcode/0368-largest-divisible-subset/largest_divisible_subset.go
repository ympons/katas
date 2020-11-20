package leetcode

import "sort"

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	sort.Ints(nums)

	D, P := make([]int, n), make([]int, n)
	max, index := 0, -1
	for i := 0; i < n; i++ {
		D[i], P[i] = 1, -1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if D[i] < D[j]+1 {
					D[i], P[i] = D[j]+1, j
				}
			}
		}
		if D[i] > max {
			max, index = D[i], i
		}
	}
	answer := []int{}
	for index != -1 {
		answer = append(answer, nums[index])
		index = P[index]
	}
	return answer
}
