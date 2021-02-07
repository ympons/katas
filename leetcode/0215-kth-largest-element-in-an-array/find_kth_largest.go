package leetcode

import (
	"container/heap"
	"sort"
)

// Time: O(N * log N)
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// Time: O(N * log K)
func findKthLargestHeap(nums []int, k int) int {
	h := append(intHeap{}, nums[:k]...)
	heap.Init(&h)
	for _, v := range nums[k:] {
		heap.Push(&h, v)
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}

// IntHeap represents a min-int heap
type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
