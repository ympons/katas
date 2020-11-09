package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minDistance(test.word1, test.word2)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	word1, word2 string
	expected     int
}{
	{
		word1:    "horse",
		word2:    "ros",
		expected: 3,
	},
	{
		word1:    "intention",
		word2:    "execution",
		expected: 5,
	},
	{
		word1:    "capricornio",
		word2:    "caricorio",
		expected: 2,
	},
}
