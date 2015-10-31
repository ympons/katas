package main

import (
	"os"
	"testing"
)

func TestNewGraph(t *testing.T) {
	f, _ := os.Open("tinyG.txt")
	g := NewGraphFromStream(f)
	t.Log(g.V(), g.E())
}
