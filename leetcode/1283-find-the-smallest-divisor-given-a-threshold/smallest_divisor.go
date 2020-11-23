package leetcode

func sum(nums []int, d int) int {
	s := 0
	for _, x := range nums {
		s += x / d
		if x%d != 0 {
			s++
		}
	}
	return s
}

func smallestDivisor(nums []int, threshold int) int {
	right := -1
	for _, x := range nums {
		if x > right {
			right = x
		}
	}
	left := 1
	for left <= right {
		middle := left + (right-left)/2
		if sum(nums, middle) > threshold {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}
