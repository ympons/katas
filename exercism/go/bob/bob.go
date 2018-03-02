package bob

import "strings"

func isYell(remark string) bool {
	upper, lower := strings.ToUpper(remark), strings.ToLower(remark)
	return remark == upper && upper != lower
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isEmpty(remark string) bool {
	return len(remark) == 0
}

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if yell, question := isYell(remark), isQuestion(remark); yell && question {
		return "Calm down, I know what I'm doing!"
	} else if yell {
		return "Whoa, chill out!"
	} else if question {
		return "Sure."
	} else if isEmpty(remark) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
