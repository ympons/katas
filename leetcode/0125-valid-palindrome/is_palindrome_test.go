package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, isPalindrome(test.input))
	}
}

var tests = []struct {
	input    string
	expected bool
}{
	{
		input:    "A man, a plan, a canal: Panama",
		expected: true,
	},
	{
		input:    "race a car",
		expected: false,
	},
	{
		input:    " 11  22......1-1",
		expected: true,
	},
}
