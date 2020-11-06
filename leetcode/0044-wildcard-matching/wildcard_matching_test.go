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
		p:        "*",
		expected: true,
	},
	{
		s:        "cb",
		p:        "?a",
		expected: false,
	},
	{
		s:        "adceb",
		p:        "*a*b",
		expected: true,
	},
	{
		s:        "acdcb",
		p:        "a*c?b",
		expected: false,
	},
}
