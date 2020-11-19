package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := longestCommonSubsequence(test.text1, test.text2)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	text1, text2 string
	expected     int
}{
	{
		text1:    "abcde",
		text2:    "ace",
		expected: 3,
	},
	{
		text1:    "abc",
		text2:    "abc",
		expected: 3,
	},
	{
		text1:    "abc",
		text2:    "def",
		expected: 0,
	},
}
