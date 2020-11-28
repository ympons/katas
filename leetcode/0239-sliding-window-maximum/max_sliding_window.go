package leetcode

import "container/heap"

func maxSlidingWindow(nums []int, k int) []int {
	pq, i := make(PQ, 0, k), 0
	for ; i < k; i++ {
		heap.Push(&pq, &Item{value: i, priority: nums[i], index: i})
	}
	window, n := []int{}, len(nums)
	for j := 0; j <= n-k; j++ {
		max := pq[0].value
		for max < j {
			heap.Pop(&pq)
			max = pq[0].value
		}
		window = append(window, nums[max])
		if i < n {
			heap.Push(&pq, &Item{value: i, priority: nums[i], index: i})
			i++
		}
	}
	return window
}

type Item struct {
	value    int
	priority int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].priority > pq[j].priority }
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
