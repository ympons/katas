package leetcode

func calculate(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	stack := make([]int, 0, 100)
	num, sign := 0, '+'
	for i := 0; i < n; i++ {
		if '0' <= s[i] && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}
		if s[i] == '*' || s[i] == '/' || s[i] == '+' || s[i] == '-' || i == n-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			default:
				top := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				if sign == '*' {
					stack = append(stack, top*num)
				} else {
					stack = append(stack, top/num)
				}
			}
			num, sign = 0, rune(s[i])
		}
	}
	sum := 0
	for _, x := range stack {
		sum += x
	}
	return sum
}
