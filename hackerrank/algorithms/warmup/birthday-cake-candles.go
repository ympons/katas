package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var in io.Reader = bufio.NewReader(os.Stdin)

func main() {
	var n, x, max, c int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x > max {
			c, max = 1, x
		} else if x == max {
			c++
		}
	}
	fmt.Println(c)
}
