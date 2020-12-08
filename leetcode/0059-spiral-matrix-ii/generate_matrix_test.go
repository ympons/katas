package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMatrix(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := generateMatrix(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    int
	expected [][]int
}{
	{
		input:    1,
		expected: [][]int{{1}},
	},
	{
		input:    2,
		expected: [][]int{{1, 2}, {4, 3}},
	},
	{
		input:    3,
		expected: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
	},
	{
		input:    4,
		expected: [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
	},
}
