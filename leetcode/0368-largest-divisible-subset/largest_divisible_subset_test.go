package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestDivisibleSubset(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := largestDivisibleSubset(test.input)
		assert.ElementsMatch(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected []int
}{
	{
		input:    []int{1, 2, 3},
		expected: []int{1, 2},
	},
	{
		input:    []int{1, 2, 4, 8},
		expected: []int{1, 2, 4, 8},
	},
	{
		input:    []int{1},
		expected: []int{1},
	},
	{
		input:    []int{3, 4, 16, 8},
		expected: []int{4, 8, 16},
	},
	{
		input:    []int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720},
		expected: []int{9, 18, 90, 180, 360, 720},
	},
}
