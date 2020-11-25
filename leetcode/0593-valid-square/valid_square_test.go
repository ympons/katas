package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidSquare(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := validSquare(test.p1, test.p2, test.p3, test.p4)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	p1, p2, p3, p4 []int
	expected       bool
}{
	{
		p1:       []int{0, 0},
		p2:       []int{1, 1},
		p3:       []int{1, 0},
		p4:       []int{0, 1},
		expected: true,
	},
	{
		p1:       []int{0, 0},
		p2:       []int{1, 1},
		p3:       []int{1, 0},
		p4:       []int{0, 12},
		expected: false,
	},
	{
		p1:       []int{0, 1},
		p2:       []int{1, 1},
		p3:       []int{1, 1},
		p4:       []int{1, 0},
		expected: false,
	},
}
