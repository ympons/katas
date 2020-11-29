package leetcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := connect(buildTree(test.input))
		assert.Equal(test.expected, output.Serialize())
	}
}

var tests = []struct {
	input    []string
	expected []string
}{
	{
		input:    []string{"1", "2", "3", "4", "5", "6", "7"},
		expected: []string{"1", "#", "2", "3", "#", "4", "5", "6", "7", "#"},
	},
	{
		input:    []string{"1", "2", "3", "4", "5", "null", "7"},
		expected: []string{"1", "#", "2", "3", "#", "4", "5", "7", "#"},
	},
}

func buildTree(nums []string) *Node {
	head := &Node{}
	return head.insertLevelOrder(nums, 0, len(nums))
}

func (root *Node) insertLevelOrder(nums []string, i, n int) *Node {
	if i < n {
		if nums[i] != "null" {
			x, _ := strconv.ParseInt(nums[i], 10, 32)
			root = &Node{Val: int(x)}
			root.Left = root.Left.insertLevelOrder(nums, 2*i+1, n)
			root.Right = root.Right.insertLevelOrder(nums, 2*i+2, n)
		}
	}
	return root
}

func (root *Node) Serialize() []string {
	output := []string{}
	q := []*Node{root}
	for len(q) > 0 {
		current := q[0]
		q = q[1:len(q)]
		if current != nil {
			output = append(output, strconv.Itoa(current.Val))
			if current.Next == nil {
				output = append(output, "#")
			}
			q = append(q, current.Left)
			q = append(q, current.Right)
		}
	}
	return output
}
