package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := sortedSquares(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input, expected []int
}{
	{
		input:    []int{-4, -1, 0, 3, 10},
		expected: []int{0, 1, 9, 16, 100},
	},
	{
		input:    []int{-7, -3, 2, 3, 11},
		expected: []int{4, 9, 9, 49, 121},
	},
}
