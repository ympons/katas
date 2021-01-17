package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountGoodRectangles(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, countGoodRectangles(test.input))
	}
}

var tests = []struct {
	input    [][]int
	expected int
}{
	{
		input:    [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}},
		expected: 3,
	},
}
