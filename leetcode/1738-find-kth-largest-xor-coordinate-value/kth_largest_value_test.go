package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargestValue(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, kthLargestValue(test.matrix, test.k))
	}
}

var tests = []struct {
	matrix   [][]int
	k        int
	expected int
}{
	{
		matrix:   [][]int{{5, 2}, {1, 6}},
		k:        1,
		expected: 7,
	},
	{
		matrix:   [][]int{{5, 2}, {1, 6}},
		k:        2,
		expected: 5,
	},
	{
		matrix:   [][]int{{5, 2}, {1, 6}},
		k:        3,
		expected: 4,
	},
	{
		matrix:   [][]int{{5, 2}, {1, 6}},
		k:        4,
		expected: 0,
	},
}
