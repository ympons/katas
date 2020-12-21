package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCherryPickup(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := cherryPickup(test.grid)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	grid     [][]int
	expected int
}{
	{
		grid:     [][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}},
		expected: 24,
	},
	{
		grid:     [][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}},
		expected: 28,
	},
	{
		grid:     [][]int{{1, 0, 0, 3}, {0, 0, 0, 3}, {0, 0, 3, 3}, {9, 0, 3, 3}},
		expected: 22,
	},
	{
		grid:     [][]int{{1, 1}, {1, 1}},
		expected: 4,
	},
}
