package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestToChar(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, shortestToChar(test.s, test.c))
	}
}

func TestShortestToChar2(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, shortestToChar2(test.s, test.c))
	}
}

var tests = []struct {
	s        string
	c        byte
	expected []int
}{
	{
		s:        "loveleetcode",
		c:        'e',
		expected: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
	},
	{
		s:        "baaa",
		c:        'b',
		expected: []int{0, 1, 2, 3},
	},
	{
		s:        "aaba",
		c:        'b',
		expected: []int{2, 1, 0, 1},
	},
	{
		s:        "abaa",
		c:        'b',
		expected: []int{1, 0, 1, 2},
	},
}
