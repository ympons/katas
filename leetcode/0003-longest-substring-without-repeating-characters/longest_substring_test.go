package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, lengthOfLongestSubstring(test.input))
	}
}

var tests = []struct {
	input    string
	expected int
}{
	{
		input:    "abcabcbb",
		expected: 3,
	},
	{
		input:    "bbbbb",
		expected: 1,
	},
	{
		input:    "pwwkew",
		expected: 3,
	},
	{
		input:    "",
		expected: 0,
	},
}
