// @ympons
package main

import (
	"bufio"
	"container/heap"
	"math"
	"os"
	"strconv"
)

type Node struct {
	id int
	w  map[*Node]int64
}

type QItem struct {
	n    *Node
	dist int64
}
type Q []QItem

func (pq Q) Len() int {
	return len(pq)
}
func (pq Q) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}
func (pq Q) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *Q) Push(i interface{}) {
	*pq = append(*pq, i.(QItem))
}
func (pq *Q) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
func (pq *Q) update(item *Node, w int64) {
	for i := 0; i < len(*pq); i++ {
		if (*pq)[i].n == item {
			(*pq)[i].dist = w
			heap.Fix(pq, i)
			break
		}
	}
}

type TGraph map[int]*Node

func (g TGraph) nod(id int) (node *Node) {
	node = g[id]
	if node == nil {
		node = &Node{
			id: id,
			w:  make(map[*Node]int64),
		}
		g[id] = node
	}
	return
}
func (g TGraph) dijkstra(x int) (dist map[*Node]int64) {
	q := &Q{}
	heap.Init(q)

	ini := g.nod(x)

	dist = make(map[*Node]int64)
	dist[ini] = 0
	for k, v := range g {
		if k != x {
			dist[v] = math.MaxInt64
		}
		item := QItem{
			n:    v,
			dist: dist[v],
		}
		heap.Push(q, item)
	}

	for q.Len() != 0 {
		u := heap.Pop(q).(QItem).n
		if dist[u] == math.MaxInt64 {
			continue
		}
		for v, w := range u.w {
			if alt := dist[u] + w; alt < dist[v] {
				dist[v] = alt
				q.update(v, alt)
			}
		}
	}
	return
}

func main() {
	o := bufio.NewWriter(os.Stdout)
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	for t, _ := strconv.Atoi(s.Text()); t > 0; t-- {
		graph := TGraph{}
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		s.Scan()
		for m, _ := strconv.Atoi(s.Text()); m > 0; m-- {
			s.Scan()
			u, _ := strconv.Atoi(s.Text())
			s.Scan()
			v, _ := strconv.Atoi(s.Text())

			unode, vnode := graph.nod(u), graph.nod(v)
			unode.w[vnode], vnode.w[unode] = int64(6), int64(6)
		}
		s.Scan()
		x, _ := strconv.Atoi(s.Text())
		dist := graph.dijkstra(x)
		for i := 1; i <= n; i++ {
			if i == x {
				continue
			}
			vi, ok := graph[i]
			if !ok {
				o.WriteString("-1 ")
				continue
			}
			md, ok := dist[vi]
			if !ok || md == math.MaxInt64 {
				o.WriteString("-1 ")
				continue
			}
			o.WriteString(strconv.FormatInt(md, 10))
			o.WriteString(" ")
		}
		o.WriteString("\n")
	}
	o.Flush()
}
