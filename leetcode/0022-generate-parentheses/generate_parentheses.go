package leetcode

var output []string

func generateParenthesis(n int) []string {
	output = []string{}
	generate(n, 0, 0, "")
	return output
}

func generate(n, left, right int, str string) {
	if left+right == n*2 {
		output = append(output, str)
	}
	if left < n {
		generate(n, left+1, right, str+"(")
	}
	if right < left {
		generate(n, left, right+1, str+")")
	}
}
