package main

import (
	"bufio"
	"fmt"
	"os"
)

func doubleMedian(count []int, d int) int {
	m, median := d/2, 0
	if d&1 == 1 {
		for i, k := 0, 0; i < 201; i++ {
			k += count[i]
			if k > m {
				median = 2 * i
				break
			}
		}
	} else {
		m1, m2 := -1, -1
		for i, k := 0, 0; i < 201; i++ {
			k += count[i]
			if m1 == -1 && k >= m {
				m1 = i
			}
			if m2 == -1 && k >= m+1 {
				m2 = i
				break
			}
		}
		median = m1 + m2
	}
	return median
}

func activityNotifications(expenditure []int, d int) int {
	count := make([]int, 201)
	for i := 0; i < d; i++ {
		count[expenditure[i]]++
	}

	notifications := 0
	for i := d; i < len(expenditure); i++ {
		if expenditure[i] >= doubleMedian(count, d) {
			notifications++
		}
		count[expenditure[i]]++
		count[expenditure[i-d]]--
	}
	return notifications
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var n, d int
	fmt.Fscan(in, &n, &d)

	expenditure := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &expenditure[i])
	}

	result := activityNotifications(expenditure, d)
	fmt.Fprintln(out, result)
	out.Flush()
}
