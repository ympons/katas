package leetcode

func longestValidParentheses(s string) int {
	n := len(s)
	D, answer := make([]int, n), 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				D[i] += 2
				if left := i - D[i]; left >= 0 {
					D[i] += D[left]
				}
			} else if left := i - D[i-1] - 1; left >= 0 && s[left] == '(' {
				D[i] += D[i-1] + 2
				if left-1 >= 0 {
					D[i] += D[left-1]
				}
			}
			if D[i] > answer {
				answer = D[i]
			}
		}
	}
	return answer
}
