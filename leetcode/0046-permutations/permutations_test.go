package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := permute(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected [][]int
}{
	{
		input:    []int{1, 2},
		expected: [][]int{{1, 2}, {2, 1}},
	},
	{
		input:    []int{1, 2, 3},
		expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}},
	},
}
