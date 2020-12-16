package leetcode

func partition(s string) [][]string {
	answer, current := [][]string{}, []string{}
	dfs(&answer, s, 0, len(s), &current)
	return answer
}

func dfs(answer *[][]string, s string, left, right int, current *[]string) {
	if left >= right {
		*answer = append(*answer, append([]string{}, *(current)...))
	}
	for i := left; i < right; i++ {
		if isPalindrome(s, left, i) {
			*current = append(*current, s[left:i+1])
			dfs(answer, s, i+1, right, current)
			*current = (*current)[:len(*current)-1]
		}
	}
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
