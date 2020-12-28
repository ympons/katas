package leetcode

func findRedundantConnection(edges [][]int) []int {
	uf := newQuickUF(1001)
	for _, edge := range edges {
		if !uf.union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}

// Weighted Quick-Union
// https://github.com/ympons/go-algorithms/blob/master/unionfind/unionfind.go#L146
type qUF struct {
	id []int
	sz []int
}

func newQuickUF(n int) *qUF {
	id, sz := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	return &qUF{id, sz}
}

func (qu *qUF) find(p int) int {
	// chase parent pointers until reach root (depth of i array accesses)
	for p != qu.id[p] {
		// this extra line is for path compression by halving
		// Keeps tree almost completely flat
		qu.id[p] = qu.id[qu.id[p]]

		p = qu.id[p]
	}
	return p
}

func (qu *qUF) union(p, q int) bool {
	i, j := qu.find(p), qu.find(q)
	if i == j {
		return false
	}

	// make root of smaller size (sz) point to root of larger size
	if qu.sz[i] < qu.sz[j] {
		qu.id[i] = j
	} else if qu.sz[i] > qu.sz[j] {
		qu.id[j] = i
	} else {
		qu.id[j] = i
		qu.sz[i]++
	}
	return true
}
