package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := nextGreaterElement(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input, expected int
}{
	{
		input:    12,
		expected: 21,
	},
	{
		input:    21,
		expected: -1,
	},
	{
		input:    231,
		expected: 312,
	},
	{
		input:    1999999999,
		expected: -1,
	},
}
