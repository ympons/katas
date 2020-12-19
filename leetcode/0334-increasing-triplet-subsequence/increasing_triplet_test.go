package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTriplet(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := increasingTriplet(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected bool
}{
	{
		input:    []int{1, 2, 3, 4, 5},
		expected: true,
	},
	{
		input:    []int{5, 4, 3, 2, 1},
		expected: false,
	},
	{
		input:    []int{2, 1, 5, 0, 4, 6},
		expected: true,
	},
}
