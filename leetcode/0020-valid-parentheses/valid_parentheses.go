package leetcode

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			n := len(stack)
			if matches := n > 0 && ((c == ')' && stack[n-1] == '(') ||
				(c == ']' && stack[n-1] == '[') ||
				(c == '}' && stack[n-1] == '{')); !matches {
				return false
			}
			stack = stack[:n-1]
		}
	}
	return len(stack) == 0
}
