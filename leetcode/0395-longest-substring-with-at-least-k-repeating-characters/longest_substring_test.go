package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstring(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := longestSubstring(test.s, test.k)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	s        string
	k        int
	expected int
}{
	{
		s:        "aaabb",
		k:        3,
		expected: 3,
	},
	{
		s:        "ababbc",
		k:        2,
		expected: 5,
	},
	{
		s:        "abcde",
		k:        1,
		expected: 5,
	},
	{
		s:        "ababacb",
		k:        3,
		expected: 0,
	},
}
