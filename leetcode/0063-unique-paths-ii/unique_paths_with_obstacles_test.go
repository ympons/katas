package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := uniquePathsWithObstacles(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    [][]int
	expected int
}{
	{
		input:    [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		expected: 2,
	},
	{
		input:    [][]int{{0, 1}, {0, 0}},
		expected: 1,
	},
}
