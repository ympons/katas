package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseStrings(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, closeStrings(test.word1, test.word2))
	}
}

var tests = []struct {
	word1, word2 string
	expected     bool
}{
	{
		word1:    "abc",
		word2:    "bca",
		expected: true,
	},
	{
		word1:    "a",
		word2:    "aa",
		expected: false,
	},
	{
		word1:    "cabbba",
		word2:    "abbccc",
		expected: true,
	},
	{
		word1:    "cabbba",
		word2:    "aabbss",
		expected: false,
	},
}
