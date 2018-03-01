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
		for t, _ := strconv.Atoi(in.Text()); t > 0 && in.Scan(); t-- {
			n, _ := strconv.ParseUint(in.Text(), 10, 32)
			r := ^uint32(n)
			out.WriteString(strconv.FormatUint(uint64(r), 10))
			out.WriteString("\n")
		}
	}
	out.Flush()
}
