package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostDP(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minCostDP(test.input)
		assert.Equal(test.expected, output)
	}
}

func TestMinCost(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minCost(test.input)
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
}
