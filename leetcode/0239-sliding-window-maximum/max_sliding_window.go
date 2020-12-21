package leetcode

import "container/heap"

// monotonic queue - O(n)
func maxSlidingWindowDeque(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	// number of windows
	window := make([]int, n-k+1)
	// store indexes
	deque := []int{}
	for i := 0; i < n; i++ {
		// remove numbers out of range k
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		// remove smaller numbers in k range
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		// add nums[i]
		deque = append(deque, i)
		// add to the windows
		if i >= k-1 {
			window[i-k+1] = nums[deque[0]]
		}
	}
	return window
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	maxLeft, maxRight := make([]int, n), make([]int, n)
	maxLeft[0] = nums[0]
	maxRight[n-1] = nums[n-1]
	for i, j := 1, n-1; i < n; i++ {
		maxLeft[i] = nums[i]
		if i%k > 0 && maxLeft[i-1] > maxLeft[i] {
			maxLeft[i] = maxLeft[i-1]
		}

		j = n - i - 1
		maxRight[j] = nums[j]
		if j%k > 0 && maxRight[j+1] > maxRight[j] {
			maxRight[j] = maxRight[j+1]
		}
	}
	window, max := make([]int, 0, n-k+1), 0
	for i := 0; i+k <= n; i++ {
		if maxRight[i] > maxLeft[i+k-1] {
			max = maxRight[i]
		} else {
			max = maxLeft[i+k-1]
		}
		window = append(window, max)
	}
	return window
}

func maxSlidingWindowPQ(nums []int, k int) []int {
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
