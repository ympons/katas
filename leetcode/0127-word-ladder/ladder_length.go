package leetcode

// Time: O(N^2 * M) N = len(word), M = len(wordList)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	WC := map[string][]string{}
	for _, word := range wordList {
		for _, w := range combo(word) {
			WC[w] = append(WC[w], word)
		}
	}

	M, answer := map[string]struct{}{}, 1
	M[beginWord] = struct{}{}
	queue := []string{beginWord}
	for len(queue) > 0 {
		for l := len(queue); l > 0; l-- {
			current := queue[0]
			queue = queue[1:]
			for _, w := range combo(current) {
				for _, word := range WC[w] {
					if _, visited := M[word]; visited {
						continue
					}
					if word == endWord {
						return answer + 1
					}
					M[word] = struct{}{}
					queue = append(queue, word)
				}
				delete(WC, w)
			}
		}
		answer++
	}
	return 0
}

func combo(a string) []string {
	n := len(a)
	W := make([]string, n)
	for i := 0; i < n; i++ {
		temp := []byte{}
		temp = append(temp, a[:i]...)
		temp = append(temp, '.')
		temp = append(temp, a[i+1:n]...)
		W[i] = string(temp)
	}
	return W
}
