package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostToMoveChips(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minCostToMoveChips(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{1, 2, 3},
		expected: 1,
	},
	{
		input:    []int{2, 2, 2, 3, 3},
		expected: 2,
	},
	{
		input:    []int{1, 1000000000},
		expected: 1,
	},
}
