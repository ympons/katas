package leetcode

func backspaceCompare(S string, T string) bool {
	return helper(S) == helper(T)
}

func helper(s string) string {
	stack := []rune{}
	for _, c := range s {
		if c == '#' {
			if n := len(stack); n > 0 {
				stack = stack[0 : n-1]
			}
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}
