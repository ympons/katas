package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaysToSplit(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := waysToSplit(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{1, 2, 2, 2, 5, 0},
		expected: 3,
	},
	{
		input:    []int{1, 1, 1},
		expected: 1,
	},
	{
		input:    []int{3, 2, 1},
		expected: 0,
	},
}
