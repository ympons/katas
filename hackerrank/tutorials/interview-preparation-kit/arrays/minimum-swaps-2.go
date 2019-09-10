package main

import (
	"fmt"
	"sort"
)

func minimumSwaps(arr []int, n int) int {
	p := make(map[int]int)
	for i, x := range arr {
		p[x] = i
	}
	sort.Ints(arr)

	visited := make(map[int]bool)
	swaps := 0
	for i := 0; i < n; i++ {
		if visited[i] || p[arr[i]] == i {
			continue
		}

		cycle := 0
		for j := i; !visited[j]; j = p[arr[j]] {
			visited[j] = true
			cycle++
		}
		swaps += cycle - 1
	}

	return swaps
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	swaps := minimumSwaps(arr, n)
	fmt.Println(swaps)
}
