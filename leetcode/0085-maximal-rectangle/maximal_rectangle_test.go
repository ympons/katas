package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalRectangle(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maximalRectangle(test.input)
		assert.Equal(test.expected, output)
	}
}

func TestMaximalRectangleDP(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maximalRectangleDP(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    [][]byte
	expected int
}{
	{
		input: [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		},
		expected: 6,
	},
	{
		input:    [][]byte{},
		expected: 0,
	},
	{
		input:    [][]byte{{'0'}},
		expected: 0,
	},
	{
		input:    [][]byte{{'1'}},
		expected: 1,
	},
	{
		input:    [][]byte{{'0', '0'}},
		expected: 0,
	},
}
