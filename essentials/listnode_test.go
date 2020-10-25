package essentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntsToList(t *testing.T) {
	assert := assert.New(t)

	input := []int{2, 3, 5, 7, 8}
	output := IntsToList(input)
	if assert.NotNil(output) {
		for i, node := 0, output; node != nil; i++ {
			assert.Equal(input[i], node.Val)
			node = node.Next
		}
	}

	assert.Nil(IntsToList([]int{}))
}

func TestListToInts(t *testing.T) {
	assert := assert.New(t)

	input := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}
	assert.Equal([]int{2, 3, 5, 7, 8}, ListToInts(input))
	assert.Equal([]int{}, ListToInts(nil))
}
