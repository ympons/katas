package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestValidParentheses(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := longestValidParentheses(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    string
	expected int
}{
	{
		input:    "(()",
		expected: 2,
	},
	{
		input:    "(())",
		expected: 4,
	},
	{
		input:    ")()())",
		expected: 4,
	},
	{
		input:    "",
		expected: 0,
	},
	{
		input:    ")(()(()))",
		expected: 8,
	},
	{
		input:    ")((()())())",
		expected: 10,
	},
}
