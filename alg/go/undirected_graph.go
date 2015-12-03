package main

import (
	"fmt"
	"io"
)

// Undirected Graph Api
type Graph interface {
	// number of vertices
	V() int
	// number of edges
	E() int
	// add an edge
	AddEdge(int, int)
	// vertices adjacent to a vertex
	Adj(int) []int
	// string representation
	String() string
}

// Just a Graph factory interface
type GraphFactory interface {
	// Create an empty graph with n vertices
	NewGraph(int) Graph
	// Create a graph from a reader
	NewGraphFromReader(io.Reader) Graph
}

type graph_ struct {
	v_, e_ int
	adj_   [][]int
}

// Create a new graph
func NewGraph(n int) *graph_ {
	adj := make([][]int, n)
	for v := 0; v < n; v++ {
		adj[v] = make([]int, 0, n)
	}
	return &graph_{v_: n, e_: 0, adj_: adj}
}

// Create a graph from a reader
func NewGraphFromReader(r io.Reader) *graph_ {
	var n, E int

	fmt.Fscanf(r, "%d", &n)
	fmt.Fscanf(r, "%d", &E)

	g := NewGraph(n)
	for i := 0; i < E; i++ {
		var v, w int
		fmt.Fscanf(r, "%d %d", &v, &w)
		g.AddEdge(v, w)
	}
	return g
}

// Add an edge v-w
func (g *graph_) AddEdge(v, w int) {
	g.adj_[v] = append(g.adj_[v], w)
	g.adj_[w] = append(g.adj_[w], v)
	g.e_++
}

// Number of vertices
func (g graph_) V() int {
	return g.v_
}

// Number of edges
func (g graph_) E() int {
	return g.e_
}

// Vertices adjacent to v
func (g graph_) Adj(v int) []int {
	return g.adj_[v]
}

// String representation
func (g graph_) String() string {
	return fmt.Sprintf("graph: V: %d E: %d -> %+v", g.V(), g.E(), g.adj_)
}

// Graph-processing:
// compute the degree of v
func degree(G Graph, v int) int {
	degree := 0
	for _ = range G.Adj(v) {
		degree++
	}
	return degree
}

// compute maximum degree
func maxDegree(G Graph) int {
	max := 0
	for v := 0; v < G.V(); v++ {
		if d := degree(G, v); d > max {
			max = d
		}
	}
	return max
}

// compute average degree
func avgDegree(G Graph) int {
	return 2 * G.E() / G.V()
}

// count self-loops
func numberOfSelfLoops(G Graph) int {
	count := 0
	for v := 0; v < G.V(); v++ {
		for _, w := range G.Adj(v) {
			if v == w {
				count++
			}
		}
	}
	return count / 2
}

// Graph processing
type Paths interface {
	HasPathTo(int) bool
	PathTo(int) []int
}

// Depth-first search
type dfsPaths_ struct {
	marked []bool
	edgeTo []int
	s      int
}

func NewDFSPaths(g Graph, s int) *dfsPaths_ {
	paths := &dfsPaths_{
		marked: make([]bool, g.V()),
		edgeTo: make([]int, g.V()),
		s:      s,
	}
	paths.dfs(g, s)

	return paths
}

func (p *dfsPaths_) dfs(g Graph, v int) {
	p.marked[v] = true
	for _, w := range g.Adj(v) {
		if !p.marked[w] {
			p.dfs(g, w)
			p.edgeTo[w] = v
		}
	}
}

func (p *dfsPaths_) iterativeDfs(g Graph, v int) {
	p.marked[v] = true
	s := NewStack()
	for _, w := range g.Adj(v) {
		s.Push(w)
	}
	for !s.IsEmpty() {
		w, _ := s.Pop()
		for _, u := range g.Adj(w) {
			if !p.marked[u] {
				p.marked[u] = true
				s.Push(u)
			}
		}
	}
}

func (p *dfsPaths_) HasPathTo(v int) bool {
	return p.marked[v]
}

func (p *dfsPaths_) PathTo(v int) []int {
	if !p.marked[v] {
		return nil
	}

	path := make([]int, 0, len(p.marked))
	for x := v; x != p.s; x = p.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, p.s)

	return path
}

// Breadth-first search
type bfsPaths_ struct {
	marked []bool
	edgeTo []int
	s      int
}

func NewBFSPaths(g Graph, s int) *bfsPaths_ {
	paths := &bfsPaths_{
		marked: make([]bool, g.V()),
		edgeTo: make([]int, g.V()),
		s:      s,
	}
	paths.bfs(g, s)

	return paths
}

func (p *bfsPaths_) bfs(g Graph, s int) {
	p.marked[s] = true
	q := NewQueue()
	q.Enqueue(s)
	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		for _, w := range g.Adj(v) {
			if !p.marked[w] {
				q.Enqueue(w)
				p.marked[w] = true
				p.edgeTo[w] = v
			}
		}
	}
}

func (p *bfsPaths_) HasPathTo(v int) bool {
	return p.marked[v]
}

func (p *bfsPaths_) PathTo(v int) []int {
	if !p.marked[v] {
		return nil
	}

	path := make([]int, 0, len(p.marked))
	for x := v; x != p.s; x = p.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, p.s)

	return path
}

// Connected Components
type CC interface{
	Connected(v, w int) bool
	Count() int
	Id(v int) int
}

type cc_ struct {
	marked []bool
	id []int
	count int
}

func NewCC(g Graph) *cc_ {
	cc := &cc_{
		marked: make([]bool, g.V()),
		id: make([]int, g.V()),
		count: 0,
	}
	for v := 0; v < g.V(); v++ {
		if !cc.marked[v] {
			cc.dfs(g, v)
			cc.count++
		}
	}

	return cc
}

func (c *cc_) dfs(g Graph, v int) {
	c.marked[v] = true
	c.id[v] = c.count
	for _, w := range g.Adj(v) {
		if !c.marked[w] {
			c.dfs(g, w)
		}
	}
}

func (c cc_) Connected(v, w int) bool {
	return c.id[v] == c.id[w]
}

func (c cc_) Count() int {
	return c.count
}

func (c cc_) Id(v int) int {
	return c.id[v]
}
