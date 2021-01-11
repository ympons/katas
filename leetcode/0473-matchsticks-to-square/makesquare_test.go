package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakesquare(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, makesquare(test.input))
	}
}

var tests = []struct {
	input    []int
	expected bool
}{
	{
		input:    []int{1, 1, 2, 2, 2},
		expected: true,
	},
	{
		input:    []int{3, 3, 3, 3, 4},
		expected: false,
	},
}
