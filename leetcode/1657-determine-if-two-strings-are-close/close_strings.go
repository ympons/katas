package leetcode

import "sort"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	C1, C2 := make([]int, 26), make([]int, 26)
	for i := range word1 {
		C1[word1[i]-'a']++
		C2[word2[i]-'a']++
	}
	for i := range C1 {
		if (C1[i] == 0 && C2[i] != 0) || (C2[i] == 0 && C1[i] != 0) {
			return false
		}
	}
	sort.Ints(C1)
	sort.Ints(C2)
	for i := range C1 {
		if C1[i] != C2[i] {
			return false
		}
	}
	return true
}
