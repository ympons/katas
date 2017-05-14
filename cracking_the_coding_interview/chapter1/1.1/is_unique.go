package strings

func isUnique(s string) bool {
	m := make(map[rune]bool)
	for _, v := range s {
		if ok := m[v]; ok {
			return false
		}
		m[v] = true
	}
	return true
}
