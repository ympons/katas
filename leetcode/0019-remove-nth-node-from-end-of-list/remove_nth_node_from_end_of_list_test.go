package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ympons/challenges/essentials"
)

func TestRemoveNthFromEnd(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := removeNthFromEnd(
			essentials.IntsToList(test.nums),
			test.n,
		)
		assert.Equal(test.expected, essentials.ListToInts(output))
	}
}

var tests = []struct {
	nums     []int
	n        int
	expected []int
}{
	{
		nums:     []int{1, 2, 3, 4, 5},
		n:        2,
		expected: []int{1, 2, 3, 5},
	},
	{
		nums:     []int{1},
		n:        1,
		expected: []int{},
	},
	{
		nums:     []int{1, 2},
		n:        1,
		expected: []int{1},
	},
}
