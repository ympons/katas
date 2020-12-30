package leetcode

const maxInt = (1 << 31) - 1

func nextGreaterElement(n int) int {
	x, digits := n, make([]int, 0, 10)
	for x != 0 {
		digits = append(digits, x%10)
		x /= 10
	}
	// reverse
	digits = reverse(digits)
	// next lexicographical permutation
	return nextPermutation(digits)
}

func nextPermutation(digits []int) int {
	// find non-increasing suffix
	i := len(digits) - 1
	for i > 0 && digits[i-1] >= digits[i] {
		i--
	}
	if i <= 0 {
		return -1
	}

	// find successor to pivot
	j := len(digits) - 1
	for digits[j] <= digits[i-1] {
		j--
	}
	digits[i-1], digits[j] = digits[j], digits[i-1]

	// reverse suffix
	j = len(digits) - 1
	for i < j {
		digits[i], digits[j] = digits[j], digits[i]
		i++
		j--
	}

	answer := toNumber(digits)
	if answer <= maxInt {
		return answer
	}

	return -1
}

func reverse(digits []int) []int {
	n := len(digits)
	for i := n/2 - 1; i >= 0; i-- {
		digits[i], digits[n-i-1] = digits[n-i-1], digits[i]
	}
	return digits
}

func toNumber(digits []int) int {
	n := 0
	for _, d := range digits {
		n = n*10 + d
	}
	return n
}
