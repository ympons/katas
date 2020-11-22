package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinHeightTrees(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := findMinHeightTrees(test.n, test.edges)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	n        int
	edges    [][]int
	expected []int
}{
	{
		n:        4,
		edges:    [][]int{{1, 0}, {1, 2}, {1, 3}},
		expected: []int{1},
	},
	{
		n:        6,
		edges:    [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
		expected: []int{3, 4},
	},
	{
		n:        1,
		edges:    [][]int{},
		expected: []int{0},
	},
	{
		n:        2,
		edges:    [][]int{{0, 1}},
		expected: []int{0, 1},
	},
}
