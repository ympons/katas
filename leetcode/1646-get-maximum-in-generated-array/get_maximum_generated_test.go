package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaximumGenerated(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, getMaximumGenerated(test.input))
	}
}

var tests = []struct {
	input    int
	expected int
}{
	{
		input:    6,
		expected: 3,
	},
	{
		input:    100,
		expected: 21,
	},
	{
		input:    0,
		expected: 0,
	},
	{
		input:    1,
		expected: 1,
	},
	{
		input:    7,
		expected: 3,
	},
	{
		input:    1000,
		expected: 89,
	},
}
