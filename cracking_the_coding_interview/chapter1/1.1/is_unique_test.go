package strings

import "testing"

func Test(t *testing.T) {
	var tests = []struct{
		s string
		exp bool
	}{
		{"abcdef1", true},
		{"aabcde", false},
		{"123456789", true},
	}
	for _, c := range tests {
		got := isUnique(c.s)
		if got != c.exp {
			t.Errorf("isUnique(%q) == %v, want %v", c.s, got, c.exp)
		}
	}
}
