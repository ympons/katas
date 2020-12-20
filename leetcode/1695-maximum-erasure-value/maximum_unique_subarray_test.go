package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumUniqueSubarray(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maximumUniqueSubarray(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{4, 2, 4, 5, 6},
		expected: 17,
	},
	{
		input:    []int{5, 2, 1, 2, 5, 2, 1, 2, 5},
		expected: 8,
	},
}
