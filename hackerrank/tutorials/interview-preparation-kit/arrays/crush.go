package main

import (
	"bufio"
	"fmt"
	"os"
)

type query struct {
	a, b int
	k    int64
}

func arrayManipulation(queries []query, n int) int64 {
	arr := make([]int64, n+1)
	for _, q := range queries {
		arr[q.a] += q.k
		if q.b+1 <= n {
			arr[q.b+1] -= q.k
		}
	}
	max, x := int64(0), int64(0)
	for i := 1; i <= n; i++ {
		x += arr[i]
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var n, m int
	fmt.Fscan(in, &n, &m)
	queries := make([]query, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &queries[i].a, &queries[i].b, &queries[i].k)
	}

	result := arrayManipulation(queries, n)
	fmt.Fprintln(out, result)
	out.Flush()
}
