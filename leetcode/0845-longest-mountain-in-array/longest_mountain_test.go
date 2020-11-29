package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestMountain(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := longestMountain(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{2, 1, 4, 7, 3, 2, 5},
		expected: 5,
	},
	{
		input:    []int{2, 2, 2},
		expected: 0,
	},
	{
		input:    []int{1, 2, 3, 2, 1, 0, 2, 3, 1},
		expected: 6,
	},
	{
		input:    []int{875, 884, 239, 731, 723, 685},
		expected: 4,
	},
}
