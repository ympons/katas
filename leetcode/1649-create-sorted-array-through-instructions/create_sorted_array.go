package leetcode

const module = int(1e9) + 7

// Using a Binary Indexed Tree or Fenwick Tree:
// https://cp-algorithms.com/data_structures/fenwick.html
// Time: O(N * log M), Space: O(M)
func createSortedArray(instructions []int) int {
	mx, answer := 0, 0
	for _, v := range instructions {
		if mx < v {
			mx = v
		}
	}
	tree := NewBIT(mx)
	for i, v := range instructions {
		answer += min(tree.sum(v-1), i-tree.sum(v))
		tree.add(v, 1)
	}
	return answer % module
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// BIT represents a Binary Indexed Tree (FenwickTree)
type BIT struct {
	tree []int
	n    int
}

// NewBIT creates a new Binary Indexed Tree
func NewBIT(n int) *BIT {
	return &BIT{
		tree: make([]int, n+1),
		n:    n + 1,
	}
}

func (f *BIT) add(idx, value int) {
	for ; idx < f.n; idx += idx & -idx {
		f.tree[idx] += value
	}
}

func (f *BIT) sum(idx int) int {
	ret := 0
	for ; idx > 0; idx -= idx & -idx {
		ret += f.tree[idx]
	}
	return ret
}
