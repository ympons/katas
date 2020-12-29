package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDiagonalOrder(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := findDiagonalOrder(test.matrix)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	matrix   [][]int
	expected []int
}{
	{
		matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		expected: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
	},
	{
		matrix:   [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		expected: []int{1, 2, 3, 5, 4, 6, 7, 8},
	},
	{
		matrix:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}},
		expected: []int{1, 2, 5, 6, 3, 4, 7, 8},
	},
}
