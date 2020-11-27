package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartition(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := canPartition(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected bool
}{
	{
		input:    []int{1, 5, 11, 5},
		expected: true,
	},
	{
		input:    []int{1, 2, 3, 5},
		expected: false,
	},
	{
		input:    []int{1, 1, 1, 1, 1},
		expected: false,
	},
	{
		input:    []int{2, 2, 1, 1},
		expected: true,
	},
}
