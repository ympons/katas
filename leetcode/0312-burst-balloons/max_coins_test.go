package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCoins(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maxCoins(test.input)
		assert.Equal(test.expected, output)
	}
}

func TestMaxCoinsRecursive(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maxCoinsRecursive(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{3, 1, 5, 8},
		expected: 167,
	},
}
