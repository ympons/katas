package leetcode

func expressiveWords(S string, words []string) int {
	answer := 0
	for _, word := range words {
		if len(word) <= len(S) && expressive(S, word) {
			answer++
		}
	}
	return answer
}

func expressive(S, word string) bool {
	i, j, n, m := 0, 0, len(S), len(word)
	for i < n && j < m {
		if S[i] != word[j] {
			return false
		}
		countS := 1
		for ; i+1 < n && S[i] == S[i+1]; i++ {
			countS++
		}
		countW := 1
		for ; j+1 < m && word[j] == word[j+1]; j++ {
			countW++
		}
		if countS != countW && countS < 3 || countS < countW {
			return false
		}
		i++
		j++
	}
	return i == n && j == m
}
