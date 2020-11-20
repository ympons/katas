package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalSquare(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maximalSquare(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    [][]byte
	expected int
}{
	{
		input: [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		},
		expected: 4,
	},
}
