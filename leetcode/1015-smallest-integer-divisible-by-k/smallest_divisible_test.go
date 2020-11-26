package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestRepunitDivByK(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := smallestRepunitDivByK(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input, expected int
}{
	{
		input:    1,
		expected: 1,
	},
	{
		input:    2,
		expected: -1,
	},
	{
		input:    3,
		expected: 3,
	},
	{
		input:    17,
		expected: 16,
	},
	{
		input:    11111111111,
		expected: 11,
	},
	{
		input:    5,
		expected: -1,
	},
	{
		input:    23,
		expected: 22,
	},
}
