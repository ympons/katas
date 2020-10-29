package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := longestPalindrome(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    string
	expected string
}{
	{
		input:    "babad",
		expected: "bab",
	},
	{
		input:    "cbbd",
		expected: "bb",
	},
	{
		input:    "a",
		expected: "a",
	},
	{
		input:    "ac",
		expected: "a",
	},
	{
		input:    "abababc",
		expected: "ababa",
	},
	{
		input:    "cabaabadxyzkl",
		expected: "abaaba",
	},
}
