package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := isPalindrome(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    int
	expected bool
}{
	{
		input:    121,
		expected: true,
	},
	{
		input:    10,
		expected: false,
	},
	{
		input:    -101,
		expected: false,
	},
	{
		input:    123321,
		expected: true,
	},
}
