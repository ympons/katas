package main

import (
	"bufio"
	"fmt"
	"os"
)

func freqQuery(queries [][]int) []int {
	freq := make([]int, 0, len(queries))
	f1 := make(map[int]int)
	fz := make(map[int]int)
	for _, q := range queries {
		switch q[0] {
		case 1:
			x := f1[q[1]]
			f1[q[1]]++
			fz[x]--
			if fz[x] <= 0 {
				delete(fz, x)
			}
			fz[x+1]++
		case 2:
			x := f1[q[1]]
			if x >= 1 {
				f1[q[1]]--
				fz[x]--
				if fz[x] <= 0 {
					delete(fz, x)
				}
				fz[x-1]++
			}
		default:
			z := 0
			if fz[q[1]] > 0 {
				z = 1
			}
			freq = append(freq, z)
		}
	}
	return freq
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var q int
	fmt.Fscan(in, &q)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = make([]int, 2)
		fmt.Fscan(in, &queries[i][0], &queries[i][1])
	}

	resp := freqQuery(queries)
	for i, x := range resp {
		fmt.Fprintf(out, "%d", x)
		if i != len(resp)-1 {
			fmt.Fprintf(out, "\n")
		}
	}

	out.Flush()
}
