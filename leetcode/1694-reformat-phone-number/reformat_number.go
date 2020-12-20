package leetcode

import "strings"

func reformatNumber(number string) string {
	digits := []byte{}
	for _, c := range number {
		if c != ' ' && c != '-' {
			digits = append(digits, byte(c))
		}
	}

	n, answer := len(digits), []string{}
	i := 0
	for ; i+3 < n && n-i > 4; i += 3 {
		answer = append(answer, string(digits[i:i+3]))
	}
	if n-i == 4 {
		answer = append(answer, string(digits[i:i+2]))
		answer = append(answer, string(digits[i+2:n]))
	} else {
		answer = append(answer, string(digits[i:n]))
	}

	return strings.Join(answer, "-")
}
