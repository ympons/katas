package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func maximumToys(prices []int, k int) int {
	sort.Ints(prices)
	max := 0
	for _, p := range prices {
		if k-p < 0 {
			break
		}
		k -= p
		max++
	}
	return max
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var n, k int
	fmt.Fscan(in, &n, &k)
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &prices[i])
	}

	resp := maximumToys(prices, k)
	fmt.Fprint(out, resp)

	out.Flush()
}
