package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, decode(test.encoded, test.first))
	}
}

var tests = []struct {
	encoded  []int
	first    int
	expected []int
}{
	{
		encoded:  []int{1, 2, 3},
		first:    1,
		expected: []int{1, 0, 2, 1},
	},
	{
		encoded:  []int{6, 2, 7, 3},
		first:    4,
		expected: []int{4, 2, 0, 7, 4},
	},
}
