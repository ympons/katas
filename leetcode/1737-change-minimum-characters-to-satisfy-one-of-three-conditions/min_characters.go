package leetcode

func minCharacters(a string, b string) int {
	answer := operation3(a, b)
	if op1 := operation12(a, b); answer > op1 {
		answer = op1
	}
	if op2 := operation12(b, a); answer > op2 {
		answer = op2
	}
	return answer
}

// cost of making every letter in a is strictly less than every letter in b
func operation12(a, b string) int {
	cost := -1
	// make a[i] <= l and b[i] > l
	for l := 'a'; l < 'z'; l++ {
		changes := 0
		// count every letter c > l in a
		for _, c := range a {
			if c > l {
				changes++
			}
		}
		// count every letter c <= l in b
		for _, c := range b {
			if c <= l {
				changes++
			}
		}
		if cost < 0 {
			cost = changes
		} else if cost > changes {
			cost = changes
		}
	}
	return cost
}

// cost of changing every letter in a and b to the same letter
func operation3(a, b string) int {
	cost := -1
	// make a[i] == l and b[i] == l
	for l := 'a'; l <= 'z'; l++ {
		changes := 0
		for _, c := range a {
			if c != l {
				changes++
			}
		}
		for _, c := range b {
			if c != l {
				changes++
			}
		}
		if cost < 0 {
			cost = changes
		} else if cost > changes {
			cost = changes
		}
	}
	return cost
}
