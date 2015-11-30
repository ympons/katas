package main

// Checking for cycles
type Cycle interface {
	HasCycle() bool
}

type cycle_ struct {
	marked   []bool
	hasCycle bool
}

func New(g Graph) *cycle_ {
	c := &cycle_{marked: make([]bool, g.V())}
	for s := 0; s < g.V(); s++ {
		if !c.marked[s] {
			c.dfs(g, s, s)
		}
	}
	return c
}

func (c *cycle_) dfs(g Graph, v, u int) {
	c.marked[v] = true
	for _, w := range g.Adj(v) {
		if !c.marked[w] {
			c.dfs(g, w, v)
		} else if w != u {
			c.hasCycle = true
		}
	}
}

func (c cycle_) HasCycle() bool {
	return c.hasCycle
}

// Bipartite graph
type BiPartite interface {
	IsBipartite() bool
}

type bipartite_ struct {
	marked      []bool
	color       []bool
	isBipartite bool
}

func NewBipartite(g Graph) *bipartite_ {
	b := &bipartite_{
		marked: make([]bool, g.V()),
		color: make([]bool, g.V()),
		isBipartite: true,
	}
	for s := 0; s < g.V(); s++ {
		if !b.marked[s] {
			b.dfs(g, s)
		}
	}

	return b
}

func (b *bipartite_) dfs(g Graph, v int) {
	b.marked[v] = true
	for _, w := range g.Adj(v) {
		if !b.marked[w] {
			b.color[w] = !b.color[v]
			b.dfs(g, w)
		} else if b.color[w] == b.color[v] {
			b.isBipartite = false
		}
	}
}

func (b bipartite_) IsBipartite() bool {
	return b.isBipartite
}

// TODO Eulerian path
type EulerianCycle interface{
	EulerianPath() []int
}

type eulerian_  struct{}

func NewEulerianPath(g Graph) *eulerian_ {
	e := &eulerian_{}
	return e
}

func (e *eulerian_) dfs(g Graph, v int) {
}
