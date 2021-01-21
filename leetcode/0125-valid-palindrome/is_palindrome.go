package leetcode

import "unicode"

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		ri, rj := rune(s[i]), rune(s[j])
		if !unicode.IsLetter(ri) && !unicode.IsNumber(ri) {
			i++
		} else if !unicode.IsLetter(rj) && !unicode.IsNumber(rj) {
			j--
		} else {
			if unicode.ToLower(ri) != unicode.ToLower(rj) {
				return false
			}
			i++
			j--
		}
	}
	return true
}
