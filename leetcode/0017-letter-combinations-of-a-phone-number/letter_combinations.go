package leetcode

var (
	phone = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	combinations []string
)

func letterCombinations(digits string) []string {
	combinations = []string{}
	if len(digits) == 0 {
		return combinations
	}

	dfs(digits, "")

	return combinations
}

func dfs(digits, str string) {
	if len(digits) == 0 {
		combinations = append(combinations, str)
		return
	}
	letters := phone[digits[0:1]]
	for _, letter := range letters {
		dfs(digits[1:], str+string(letter))
	}
}
