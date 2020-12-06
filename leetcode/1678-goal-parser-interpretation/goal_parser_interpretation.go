package leetcode

import "strings"

func interpret(command string) string {
	n := len(command)
	var parser strings.Builder
	for i := 0; i < n; {
		if command[i] == 'G' {
			parser.WriteByte(command[i])
		} else if command[i] == '(' {
			if command[i+1] == ')' {
				parser.WriteRune('o')
				i++
			} else {
				parser.WriteString("al")
				i += 2
			}
		}
		i++
	}
	return parser.String()
}
