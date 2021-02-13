package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, shortestPathBinaryMatrix((test.input)))
	}
}

var tests = []struct {
	input    [][]int
	expected int
}{
	{
		input:    [][]int{{0, 1}, {1, 0}},
		expected: 2,
	},
	{
		input:    [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
		expected: 4,
	},
	{
		input:    [][]int{{1, 0}, {1, 0}},
		expected: -1,
	},
	{
		input:    [][]int{{0, 1}, {1, 1}},
		expected: -1,
	},
}
