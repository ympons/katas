package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSmallestString(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, getSmallestString(test.n, test.k))
	}
}

var tests = []struct {
	n, k     int
	expected string
}{
	{
		n:        3,
		k:        27,
		expected: "aay",
	},
	{
		n:        5,
		k:        73,
		expected: "aaszz",
	},
}
