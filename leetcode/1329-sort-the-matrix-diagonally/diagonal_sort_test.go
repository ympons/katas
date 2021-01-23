package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalSort(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, diagonalSort(test.input))
	}
}

var tests = []struct {
	input, expected [][]int
}{
	{
		input:    [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}},
		expected: [][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}},
	},
}
