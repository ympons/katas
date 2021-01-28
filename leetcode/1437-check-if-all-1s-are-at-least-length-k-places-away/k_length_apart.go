package leetcode

func kLengthApart(nums []int, k int) bool {
	prev := -1
	for i, v := range nums {
		if v == 1 {
			if prev != -1 && i-prev <= k {
				return false
			}
			prev = i
		}
	}
	return true
}
