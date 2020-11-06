package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := isMatch(test.s, test.p)
		assert.Equal(test.expected, output)
	}
}

func TestIsMatchRecursive(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := isMatchRecursive(test.s, test.p)
		assert.Equal(test.expected, output)
	}
}

func TestIsMatchDP(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := isMatchDP(test.s, test.p)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	s, p     string
	expected bool
}{
	{
		s:        "aa",
		p:        "a",
		expected: false,
	},
	{
		s:        "aa",
		p:        "a*",
		expected: true,
	},
	{
		s:        "ab",
		p:        ".*",
		expected: true,
	},
	{
		s:        "aab",
		p:        "c*a*b",
		expected: true,
	},
	{
		s:        "mississippi",
		p:        "mis*is*p*.",
		expected: false,
	},
}
