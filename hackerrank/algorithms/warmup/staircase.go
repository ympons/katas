package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	if in.Scan() {
		n, _ := strconv.Atoi(in.Text())
		for i := 1; i <= n; i++ {
			for j := 0; j < n - i; j++ {
				out.WriteString(" ")
			}
			for k := 0; k < i; k++ {
				out.WriteString("#")
			}
			out.WriteString("\n")
		}
	}
	out.Flush()
}
