package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := generateParenthesis(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    int
	expected []string
}{
	{
		input:    1,
		expected: []string{"()"},
	},
	{
		input:    3,
		expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
	},
	{
		input:    2,
		expected: []string{"(())", "()()"},
	},
}
