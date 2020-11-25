package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlipAndInvertImage(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := flipAndInvertImage(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    [][]int
	expected [][]int
}{
	{
		input:    [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
		expected: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
	},
}
