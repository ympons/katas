package leetcode

import "container/heap"

// Time: O(N * log N)
func eatenApples(apples []int, days []int) int {
	// This is a greedy problem. We can track the available apples and eat the ones closest to the expiration date.
	answer, n := 0, len(apples)

	pq := make(PQ, 0, n)
	for i := 0; i < n || pq.Len() > 0; i++ {
		// add the available apples for the day
		if i < n && apples[i] > 0 {
			heap.Push(&pq, Item{day: i + days[i], apples: apples[i]})
		}
		// get rid of the rotten apples and the items with no apples
		for pq.Len() > 0 && (pq[0].day <= i || pq[0].apples <= 0) {
			heap.Pop(&pq)
		}
		// count the apple that can be consumed
		if pq.Len() > 0 {
			answer++
			pq[0].apples--
		}
	}

	return answer
}

// Item represents an element in the Priority Queue
type Item struct {
	day    int
	apples int
}

// PQ represents a Priority Queue
type PQ []Item

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].day < pq[j].day }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

// Push adds an item to the queue
func (pq *PQ) Push(x interface{}) {
	item := x.(Item)
	*pq = append(*pq, item)
}

// Pop remove the item from the queue
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
