package leetcode

const maxInt = 1<<31 - 1
const minInt = -(maxInt + 1)

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func myAtoi(s string) int {
	n, i := len(s), 0

	// skip whitespaces
	for i < n && s[i] == ' ' {
		i++
	}

	// check the sign
	sign := 1
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}

	num, digit := 0, 0
	for ; i < n && isDigit(s[i]); i++ {
		digit = int(s[i] - '0')
		if limit := (maxInt - digit) / 10; num > limit {
			if sign == -1 {
				return minInt
			}
			return maxInt
		}
		num = num*10 + digit
	}

	return sign * num
}
