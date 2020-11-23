package leetcode

import "strings"

var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	uniq := map[string]struct{}{}
	for _, word := range words {
		var encoded strings.Builder
		n := len(word)
		for i := 0; i < n; i++ {
			encoded.WriteString(morse[word[i]-97])
		}
		uniq[encoded.String()] = struct{}{}
	}
	return len(uniq)
}
