package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var in io.Reader = bufio.NewReader(os.Stdin)

func main() {
	var n int
	fmt.Fscan(in, &n)
	m := make(map[int]int)
	var x int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		m[x]++
	}
	count := 0
	for _, v := range m {
		count += v / 2
	}
	fmt.Println(count)
}
