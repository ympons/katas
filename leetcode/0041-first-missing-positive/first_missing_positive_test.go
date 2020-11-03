package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := firstMissingPositive(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{1, 2, 0},
		expected: 3,
	},
	{
		input:    []int{3, 4, -1, 1},
		expected: 2,
	},
	{
		input:    []int{7, 8, 9, 11, 12},
		expected: 1,
	},
}
