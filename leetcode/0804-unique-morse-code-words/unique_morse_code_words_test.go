package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := uniqueMorseRepresentations(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []string
	expected int
}{
	{
		input:    []string{"gin", "zen", "gig", "msg"},
		expected: 2,
	},
}
