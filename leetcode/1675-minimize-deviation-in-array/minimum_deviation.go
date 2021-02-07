package leetcode

import "container/heap"

const maxInt = int(1e9) + 1

// Time: ~O(N * log M * log N)
// When all the numbers (N) are the max possible power of 2 (M), it will take N * log M operations to reduce them.
// Since every operation manipulates the heap (push/pop -> log N), the worst case will take O(N * log M * log N)
func minimumDeviation(nums []int) int {
	h, minX := make(intHeap, 0, len(nums)), maxInt
	for i := range nums {
		if nums[i]&1 == 1 {
			nums[i] *= 2
		}
		minX = min(minX, nums[i])
		heap.Push(&h, nums[i])
	}

	deviation := maxInt
	for {
		x := heap.Pop(&h).(int)
		deviation = min(deviation, x-minX)
		if x&1 == 1 {
			break
		}
		minX = min(minX, x/2)
		heap.Push(&h, x/2)
	}
	return deviation
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// IntHeap represents a max-int heap
type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
