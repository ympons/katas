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
		np, idx := make(map[uint64]bool), make([]uint64, 0, 10001)
		var i, j uint64
		for i = 2; i <= 104743; i++ {
			if !np[i] {
				idx = append(idx, i)
				for j = i + i; j <= 104743; j += i {
					np[j] = true
				}
			}
		}
		for t, _ := strconv.Atoi(scanner.Text()); t > 0 && scanner.Scan(); t-- {
			n, _ := strconv.Atoi(scanner.Text())
			wout.WriteString(strconv.FormatUint(idx[n-1], 10))
			wout.WriteString("\n")
		}
		wout.Flush()
	}
}
