package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFormArray(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := canFormArray(test.arr, test.pieces)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	arr      []int
	pieces   [][]int
	expected bool
}{
	{
		arr:      []int{80},
		pieces:   [][]int{{80}},
		expected: true,
	},
	{
		arr:      []int{10, 80},
		pieces:   [][]int{{10}, {80}},
		expected: true,
	},
	{
		arr:      []int{49, 18, 16},
		pieces:   [][]int{{16, 18, 49}},
		expected: false,
	},
	{
		arr:      []int{91, 4, 64, 78},
		pieces:   [][]int{{78}, {4, 64}, {91}},
		expected: true,
	},
	{
		arr:      []int{1, 3, 5, 7},
		pieces:   [][]int{{2, 4, 6, 8}},
		expected: false,
	},
}
