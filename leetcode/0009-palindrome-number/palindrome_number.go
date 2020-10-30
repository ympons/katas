package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	digits := make([]int, 0, 10)
	for x != 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	n := len(digits)
	for i := 0; i < n/2; i++ {
		if digits[i] != digits[n-i-1] {
			return false
		}
	}

	return true
}
