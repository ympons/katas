package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := trap(test.input)
		assert.Equal(test.expected, output)
	}
}

func TestTrapDP(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := trapDP(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		expected: 6,
	},
	{
		input:    []int{4, 2, 0, 3, 2, 5},
		expected: 9,
	},
}
