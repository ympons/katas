package leetcode

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	G, V := map[int][]int{}, map[int]struct{}{}
	for _, edge := range allowedSwaps {
		G[edge[0]] = append(G[edge[0]], edge[1])
		G[edge[1]] = append(G[edge[1]], edge[0])
	}

	n := len(source)
	distance, M := n, map[int]int{}

	var dfs func(int)
	dfs = func(p int) {
		V[p] = struct{}{}

		M[source[p]]++
		if M[source[p]] <= 0 {
			distance--
		}

		M[target[p]]--
		if M[target[p]] >= 0 {
			distance--
		}

		for _, v := range G[p] {
			if _, seen := V[v]; seen {
				continue
			}
			dfs(v)
		}
	}

	for i := 0; i < n; i++ {
		if _, seen := V[i]; seen {
			continue
		}
		M = map[int]int{}
		dfs(i)
	}

	return distance
}

// Using Union Find
func minimumHammingDistanceUF(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	uf := newUF(n)
	for _, edge := range allowedSwaps {
		uf.union(edge[0], edge[1])
	}

	G := map[int][]int{}
	for i := 0; i < n; i++ {
		parent := uf.find(i)
		G[parent] = append(G[parent], i)
	}

	distance, M := n, map[int]int{}
	for _, nodes := range G {
		M = map[int]int{}
		for _, v := range nodes {
			M[source[v]]++
			if M[source[v]] <= 0 {
				distance--
			}
			M[target[v]]--
			if M[target[v]] >= 0 {
				distance--
			}
		}
	}

	return distance
}

// UF is a Weighted Quick-Union
// https://github.com/ympons/go-algorithms/blob/master/unionfind/unionfind.go#L146
type UF struct {
	parent, rank []int
}

func newUF(n int) *UF {
	parent, rank := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UF{parent, rank}
}

func (uf *UF) find(p int) int {
	if p != uf.parent[p] {
		// path compression
		uf.parent[p] = uf.find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *UF) union(p, q int) bool {
	i, j := uf.find(p), uf.find(q)
	// check if they're already connected
	if i == j {
		return false
	}
	// make root of smaller size (sz) point to root of larger size
	if uf.rank[i] < uf.rank[j] {
		uf.parent[i] = j
	} else if uf.rank[i] > uf.rank[j] {
		uf.parent[j] = i
	} else {
		uf.parent[j] = i
		uf.rank[i]++
	}
	return true
}
