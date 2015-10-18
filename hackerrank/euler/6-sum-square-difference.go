package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wout := bufio.NewWriter(os.Stdout)
	if scanner.Scan() {
		for t, _ := strconv.Atoi(scanner.Text()); t > 0 && scanner.Scan(); t-- {
			n, _ := strconv.Atoi(scanner.Text())
			s := uint64(n * (n*n - 1) * (3*n + 2) / 12)
			wout.WriteString(strconv.FormatUint(s, 10))
			wout.WriteString("\n")
		}
		wout.Flush()
	}
}
