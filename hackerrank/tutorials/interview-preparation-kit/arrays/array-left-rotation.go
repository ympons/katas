package main

import (
	"bufio"
	"fmt"
	"os"
)

func rotateArray(arr []int, d int) []int {
	i := d % len(arr)
	return append(arr[i:], arr[:i]...)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, d int
	fmt.Fscan(in, &n, &d)

	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i])
	}

	rotated := rotateArray(l, d)
	for _, x := range rotated {
		fmt.Printf("%d ", x)
	}
}
