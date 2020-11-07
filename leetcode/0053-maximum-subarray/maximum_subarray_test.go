package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maxSubArray(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		expected: 6,
	},
	{
		input:    []int{1},
		expected: 1,
	},
	{
		input:    []int{-1},
		expected: -1,
	},
	{
		input:    []int{-2147483647},
		expected: -2147483647,
	},
}
