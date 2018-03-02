package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(input string) string {
	prev := ' '
	return strings.Map(
		func(r rune) rune {
			if (unicode.IsSpace(prev) || prev == '-') && unicode.IsLetter(r) {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return -1
		},
		input)
}
