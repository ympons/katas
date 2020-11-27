package leetcode

import "strings"

func longestSubstring(s string, k int) int {
	n := len(s)
	if n < k {
		return 0
	}

	counter := make([]int, 26)
	for i := 0; i < n; i++ {
		counter[s[i]-'a']++
	}

	split := byte(0)
	for i, c := range counter {
		if 0 < c && c < k {
			split = byte(i + 'a')
			break
		}
	}

	if split == 0 {
		return n
	}

	answer := 0
	for _, str := range strings.Split(s, string(split)) {
		if tmp := longestSubstring(str, k); tmp > answer {
			answer = tmp
		}
	}
	return answer
}
