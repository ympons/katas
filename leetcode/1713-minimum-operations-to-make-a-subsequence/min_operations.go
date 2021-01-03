package leetcode

import "sort"

func minOperations(target []int, arr []int) int {
	H := map[int]int{}
	for i, v := range target {
		H[v] = i
	}

	S := make([]int, 0, len(target))
	for _, x := range arr {
		idx, exists := H[x]
		if exists {
			S = append(S, idx)
		}
	}

	return len(target) - lengthOfLIS(S)
}

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
