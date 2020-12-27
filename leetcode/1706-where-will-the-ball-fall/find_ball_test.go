package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBall(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := findBall(test.grid)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	grid     [][]int
	expected []int
}{
	{
		grid:     [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}},
		expected: []int{1, -1, -1, -1, -1},
	},
	{
		grid:     [][]int{{-1}},
		expected: []int{-1},
	},
	{
		grid:     [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, 1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, -1, -1}},
		expected: []int{2, 3, -1, -1, -1},
	},
	{
		grid:     [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, 1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, -1, 1}},
		expected: []int{2, -1, -1, -1, -1},
	},
}
