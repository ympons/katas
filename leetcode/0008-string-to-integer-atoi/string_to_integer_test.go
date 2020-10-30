package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := myAtoi(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    string
	expected int
}{
	{
		input:    "42",
		expected: 42,
	},
	{
		input:    "   -42",
		expected: -42,
	},
	{
		input:    "4193 with words",
		expected: 4193,
	},
	{
		input:    "words and 987",
		expected: 0,
	},
	{
		input:    "-91283472332",
		expected: -2147483648,
	},
	{
		input:    "-2147483648",
		expected: -2147483648,
	},
	{
		input:    "-100000000000000000000000000",
		expected: -2147483648,
	},
	{
		input:    "100000000000000000000000000",
		expected: 2147483647,
	},
}
