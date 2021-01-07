package leetcode

// O(log n)
func findKthPositiveLog(arr []int, k int) int {
	// Given A = [2, 4, 5, 8], we can create B as the missing numbers count for every position
	// B[i] = A[i] - i - 1, e.g. B = [2 - 1, 4 - 2, 5 - 3, 8 - 4] = [1, 2, 2, 4]
	// Find ith where B[i] >= k (equivalent to A[i] - i - 1 >= k)
	//  -> answer = A[ith] - (B[ith] - k + 1),  B[ith] = A[ith] - ith - 1
	//  -> answer = A[ith] - (A[ith] - ith - 1 - k + 1) = ith + k
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid]-mid-1 < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right + k
}

// O(n)
func findKthPositive(arr []int, k int) int {
	i, n := 0, 0
	for k > 0 {
		n++
		if i < len(arr) && arr[i] == n {
			i++
		} else {
			k--
		}
	}
	return n
}
