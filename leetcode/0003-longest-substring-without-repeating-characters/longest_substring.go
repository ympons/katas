package leetcode

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
