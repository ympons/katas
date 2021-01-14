package leetcode

// Time: O(n), Space: O(1) -- Sliding window
// Find the longest subarray whose sum = total_sum - x. Return n - len(subarray)
func minOperations(nums []int, x int) int {
	target, n := 0, len(nums)
	for _, v := range nums {
		target += v
	}
	target -= x
	if target == 0 {
		return n
	}
	if target < 0 {
		return -1
	}

	current, i, size := 0, -1, -1
	for j, v := range nums {
		current += v
		for i+1 < n && current > target {
			i++
			current -= nums[i]
		}
		if current == target && size < j-i {
			size = j - i
		}
	}
	if size < 0 {
		return -1
	}
	return n - size
}

// Time: O(n), Space: O(n) -- Using a map, similar to the 1. Two sum problem
func minOperationsMap(nums []int, x int) int {
	n, sum := len(nums), 0
	for _, v := range nums {
		sum += v
	}
	target := sum - x
	if target == 0 {
		return n
	}
	if target < 0 {
		return -1
	}

	current, size, M := 0, -1, map[int]int{0: -1}
	for j, v := range nums {
		current += v
		if i, ok := M[current-target]; ok && size < j-i {
			size = j - i
		}
		M[current] = j
	}
	if size < 0 {
		return -1
	}
	return n - size
}
