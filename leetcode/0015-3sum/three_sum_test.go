package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := threeSum(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected [][]int
}{
	{
		input:    []int{-1, 0, 1, 2, -1, -4},
		expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		input:    []int{},
		expected: [][]int{},
	},
	{
		input:    []int{0},
		expected: [][]int{},
	},
}
