package leetcode

const maxInt = 1 << 31

func increasingTriplet(nums []int) bool {
	n1, n2 := maxInt, maxInt
	for _, n := range nums {
		if n <= n1 {
			n1 = n
		} else if n <= n2 {
			n2 = n
		} else {
			return true
		}
	}
	return false
}
