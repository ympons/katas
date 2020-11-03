package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ympons/katas/essentials"
)

func TestMergeKLists(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		list := []*essentials.ListNode{}
		for _, ints := range test.input {
			list = append(list, essentials.IntsToList(ints))
		}
		output := mergeKLists(list)
		assert.Equal(test.expected, essentials.ListToInts(output))
	}
}

var tests = []struct {
	input    [][]int
	expected []int
}{
	{
		input:    [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
		expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
	},
	{
		input:    [][]int{},
		expected: []int{},
	},
	{
		input:    [][]int{{}},
		expected: []int{},
	},
}
