package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantConnection(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := findRedundantConnection(test.edges)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	edges    [][]int
	expected []int
}{
	{
		edges:    [][]int{{1, 2}, {1, 3}, {2, 3}},
		expected: []int{2, 3},
	},
	{
		edges:    [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
		expected: []int{1, 4},
	},
}
