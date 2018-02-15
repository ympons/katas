func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	resp := 0
	for i, j := 0, 0; j < len(s); j++ {
		if v, ok := m[s[j]]; ok {
			i = max(v, i)
		}
		resp = max(resp, j-i+1)
		m[s[j]] = j + 1
	}
	return resp
}

/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/