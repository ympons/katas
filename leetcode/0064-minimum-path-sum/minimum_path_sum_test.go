package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minPathSum(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    [][]int
	expected int
}{
	{
		input:    [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
		expected: 7,
	},
	{
		input:    [][]int{{1, 2, 3}, {4, 5, 6}},
		expected: 12,
	},
}
