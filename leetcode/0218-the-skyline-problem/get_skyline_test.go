package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSkyline(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		ouput := getSkyline(test.input)
		assert.Equal(test.expected, ouput)
	}
}

var tests = []struct {
	input    [][]int
	expected [][]int
}{
	{
		input:    [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
		expected: [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
	},
}
