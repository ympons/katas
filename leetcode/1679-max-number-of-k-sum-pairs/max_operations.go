package leetcode

import "sort"

// Time: O(N * log N)
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

// Time: O(N), Space: O(N)
func maxOperationsMap(nums []int, k int) int {
	M, answer := map[int]int{}, 0
	for _, v := range nums {
		if M[k-v] > 0 {
			answer++
			M[k-v]--
			if M[k-v] <= 0 {
				delete(M, k-v)
			}
		} else {
			M[v]++
		}
	}
	return answer
}
