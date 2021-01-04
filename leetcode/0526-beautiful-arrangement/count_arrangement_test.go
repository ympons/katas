package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountArrangementBitMask(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := countArrangementBitMask(test.input)
		assert.Equal(test.expected, output)
	}
}

func TestCountArrangement(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := countArrangement(test.input)
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
		expected: 2,
	},
	{
		input:    3,
		expected: 3,
	},
	{
		input:    4,
		expected: 8,
	},
	{
		input:    5,
		expected: 10,
	},
	{
		input:    15,
		expected: 24679,
	},
}
