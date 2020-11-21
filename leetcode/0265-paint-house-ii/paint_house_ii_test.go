package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostII(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minCostII(test.input)
		assert.Equal(test.expected, output)
	}
}

func TestMinCostIIOptimized(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minCostIIOptimized(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    [][]int
	expected int
}{
	{
		input:    [][]int{},
		expected: 0,
	},
	{
		input:    [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		expected: 3,
	},
	{
		input:    [][]int{{10, 3, 5}, {2, 1, 6}, {5, 6, 8}},
		expected: 11,
	},
	{
		input:    [][]int{{5, 2}, {4, 2}},
		expected: 6,
	},
}
