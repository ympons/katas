package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ympons/katas/essentials"
)

func TestSwapPairs(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := swapPairs(essentials.IntsToList(test.input))
		assert.Equal(test.expected, essentials.ListToInts(output))
	}
}

var tests = []struct {
	input    []int
	expected []int
}{
	{
		input:    []int{1, 2, 3, 4},
		expected: []int{2, 1, 4, 3},
	},
	{
		input:    []int{},
		expected: []int{},
	},
	{
		input:    []int{1},
		expected: []int{1},
	},
}
