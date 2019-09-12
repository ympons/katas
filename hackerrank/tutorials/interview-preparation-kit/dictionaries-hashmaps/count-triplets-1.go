package main

import (
	"bufio"
	"fmt"
	"os"
)

func countTriplets(arr []int64, r int64) int64 {
	h2 := make(map[int64]int64)
	h3 := make(map[int64]int64)

	triples := int64(0)
	for _, a := range arr {
		triples += h3[a]
		h3[a*r] += h2[a]
		h2[a*r]++
	}
	return triples
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var (
		n int
		r int64
	)
	fmt.Fscan(in, &n, &r)
	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	result := countTriplets(arr, r)
	fmt.Fprintln(out, result)
	out.Flush()
}
