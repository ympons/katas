package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleArea(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := largestRectangleArea(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{2, 1, 5, 6, 2, 3},
		expected: 10,
	},
	{
		input:    []int{2, 1, 1, 1, 1, 1, 2, 3},
		expected: 8,
	},
}
