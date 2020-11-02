package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := isValid(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    string
	expected bool
}{
	{
		input:    "()",
		expected: true,
	},
	{
		input:    "()[]{}",
		expected: true,
	},
	{
		input:    "(]",
		expected: false,
	},
	{
		input:    "([)]",
		expected: false,
	},
	{
		input:    "{[]}",
		expected: true,
	},
}
