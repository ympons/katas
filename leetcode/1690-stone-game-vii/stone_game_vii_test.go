package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoneGameVII(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := stoneGameVII(test.input)
		assert.Equal(test.expected, output)
	}
}

func TestStoneGameVIIRecursive(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := stoneGameVIIRecursive(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{5, 3, 1, 4, 2},
		expected: 6,
	},
	{
		input:    []int{7, 90, 5, 1, 100, 10, 10, 2},
		expected: 122,
	},
}
