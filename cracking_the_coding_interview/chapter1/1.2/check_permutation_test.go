package strings

import "testing"

func Test(t *testing.T) {
	var tests = []struct{
		s, v string
		exp bool
	}{
		{"abcdefg", "abc", true},
		{"123456", "abcd", false},
	}
	for _, c := range tests {
		got := CheckPermutation(c.s, c.v)
		if got != c.exp {
			t.Errorf("CheckPermutation(%q, %q) == %v, want %v", c.s, c.v, got, c.exp)
		}
	}
}
