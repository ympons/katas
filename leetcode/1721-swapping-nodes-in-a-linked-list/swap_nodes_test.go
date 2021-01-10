package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ympons/katas/essentials"
)

func TestSwapNodes(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := swapNodes(
			essentials.IntsToList(test.list),
			test.k,
		)
		assert.Equal(test.expected, essentials.ListToInts(output))
	}
}

var tests = []struct {
	list     []int
	k        int
	expected []int
}{
	{
		list:     []int{1, 2, 3, 4, 5},
		k:        2,
		expected: []int{1, 4, 3, 2, 5},
	},
	{
		list:     []int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5},
		k:        5,
		expected: []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5},
	},
	{
		list:     []int{1},
		k:        1,
		expected: []int{1},
	},
	{
		list:     []int{1, 2},
		k:        1,
		expected: []int{2, 1},
	},
	{
		list:     []int{1, 2, 3},
		k:        2,
		expected: []int{1, 2, 3},
	},
}
