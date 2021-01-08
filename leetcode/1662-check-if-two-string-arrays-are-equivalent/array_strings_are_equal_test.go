package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayStringsAreEqual(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := arrayStringsAreEqual(test.word1, test.word2)
		assert.Equal(test.expected, output)
	}
}

func TestArrayStringsAreEqualBuffer(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := arrayStringsAreEqualBuffer(test.word1, test.word2)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	word1, word2 []string
	expected     bool
}{
	{
		word1:    []string{"ab", "c"},
		word2:    []string{"a", "bc"},
		expected: true,
	},
	{
		word1:    []string{"a", "cb"},
		word2:    []string{"ab", "c"},
		expected: false,
	},
	{
		word1:    []string{"abc", "d", "defg"},
		word2:    []string{"abcddefg"},
		expected: true,
	},
}
