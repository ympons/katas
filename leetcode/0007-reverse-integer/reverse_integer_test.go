package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := reverse(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input, expected int
}{
	{
		input:    123,
		expected: 321,
	},
	{
		input:    -123,
		expected: -321,
	},
	{
		input:    120,
		expected: 21,
	},
	{
		input:    0,
		expected: 0,
	},
}
