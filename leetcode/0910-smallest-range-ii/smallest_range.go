package leetcode

import "sort"

func smallestRangeII(A []int, k int) int {
	sort.Ints(A)
	n := len(A)
	max, min := A[n-1], A[0]
	answer := max - min
	for i := 0; i < n-1; i++ {
		max, min = A[i]+k, A[i+1]-k
		if max < A[n-1]-k {
			max = A[n-1] - k
		}
		if min > A[0]+k {
			min = A[0] + k
		}
		if answer > max-min {
			answer = max - min
		}
	}
	return answer
}
