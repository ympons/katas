package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairsDP(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minCostClimbingStairsDP(test.input)
		assert.Equal(test.expected, output)
	}
}

func TestMinCostClimbingStairs(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minCostClimbingStairs(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{10, 15, 20},
		expected: 15,
	},
	{
		input:    []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		expected: 6,
	},
}
