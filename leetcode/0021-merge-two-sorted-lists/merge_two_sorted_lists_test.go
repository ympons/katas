package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ympons/katas/essentials"
)

func TestMergeTwoLists(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := mergeTwoLists(
			essentials.IntsToList(test.l1),
			essentials.IntsToList(test.l2),
		)
		assert.Equal(test.expected, essentials.ListToInts(output))
	}
}

func TestMergeTwoListsRecursive(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := mergeTwoListsRecursive(
			essentials.IntsToList(test.l1),
			essentials.IntsToList(test.l2),
		)
		assert.Equal(test.expected, essentials.ListToInts(output))
	}
}

var tests = []struct {
	l1, l2   []int
	expected []int
}{
	{
		l1:       []int{1, 2, 4},
		l2:       []int{1, 3, 4},
		expected: []int{1, 1, 2, 3, 4, 4},
	},
	{
		l1:       []int{},
		l2:       []int{},
		expected: []int{},
	},
	{
		l1:       []int{},
		l2:       []int{0},
		expected: []int{0},
	},
}
