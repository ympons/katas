package main

import (
	"os"
	"testing"
)

func TestNewGraph(t *testing.T) {
	f, _ := os.Open("tinyG.txt")
	g := NewGraphFromReader(f)
	t.Log(g.V(), g.E())
}

func TestDFSPaths(t *testing.T) {
	f, _ := os.Open("tinyG.txt")
	g := NewGraphFromReader(f)
	paths := NewDFSPaths(g, 1)
	for v := 0; v < g.V(); v++ {
		if paths.HasPathTo(v) {
			t.Log(v)
		}
	}
}

func TestBFSPaths(t *testing.T) {
	f, _ := os.Open("tinyG.txt")
	g := NewGraphFromReader(f)
	paths := NewBFSPaths(g, 1)
	for v := 0; v < g.V(); v++ {
		if paths.HasPathTo(v) {
			t.Log(v)
		}
	}
}
