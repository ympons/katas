package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWinnerSquareGame(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		ouput := winnerSquareGame(test.input)
		assert.Equal(test.expected, ouput)
	}
}

var tests = []struct {
	input    int
	expected bool
}{
	{
		input:    1,
		expected: true,
	},
	{
		input:    2,
		expected: false,
	},
	{
		input:    4,
		expected: true,
	},
	{
		input:    7,
		expected: false,
	},
	{
		input:    5,
		expected: false,
	},
}
