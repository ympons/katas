package main

import (
	"os"
	"testing"
)

type Fatalistic interface{
	Fatal(...interface{})
}

func graphTest(t Fatalistic) *graph_ {
	f, err := os.Open("tinyG.txt")
	if err != nil {
		t.Fatal(err)
	}
	return NewGraphFromReader(f)
}

func TestNewGraph(t *testing.T) {
	g := graphTest(t)
	t.Log(g.V(), g.E())
}

func TestDFSPaths(t *testing.T) {
	g := graphTest(t)
	paths := NewDFSPaths(g, 1)
	for v := 0; v < g.V(); v++ {
		if paths.HasPathTo(v) {
			t.Log(v)
		}
	}
}

func TestBFSPaths(t *testing.T) {
	g := graphTest(t)
	paths := NewBFSPaths(g, 1)
	for v := 0; v < g.V(); v++ {
		if paths.HasPathTo(v) {
			t.Log(v)
		}
	}
}

func TestCC(t *testing.T) {
	g := graphTest(t)
	cc := NewCC(g)
	t.Log(cc.Count())
}
