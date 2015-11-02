package main

import (
	"io"
	"fmt"
)

// Undirected Graph Api
type Graph interface {
	// number of vertices
	V() int
	// number of edges
	E() int
	// add an edge
	AddEgde(int, int)
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
	adj_ [][]int
}

// Create a new graph
func NewGraph(n int) *graph_{
	adj := make([][]int, n)
	for v := 0; v < n; v++ {
		adj[v] = make([]int, 0, n)
	}
	return &graph_{ v_: n, e_: 0, adj_: adj }
}

// Create a graph from a reader
func NewGraphFromStream(r io.Reader) *graph_{
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
	return ""
}

// Graph-processing:
// compute the degree of v
func degree(G Graph, v int) int {
	degree := 0;
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
	return count/2
}
