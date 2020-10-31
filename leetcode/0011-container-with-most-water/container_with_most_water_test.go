package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maxArea(test.input)
		assert.Equal(test.expected, output)
	}
}

func TestMaxAreaNaive(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maxAreaNaive(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		expected: 49,
	},
	{
		input:    []int{1, 1},
		expected: 1,
	},
	{
		input:    []int{4, 3, 2, 1, 4},
		expected: 16,
	},
	{
		input:    []int{1, 2, 1},
		expected: 2,
	},
}
