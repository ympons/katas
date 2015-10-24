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
		np, sum := make(map[uint64]bool), make([]uint64, 2, 1000001)
		var i, j uint64
		for i = 2; i <= 1000000; i++ {
			if !np[i] {
				sum = append(sum, sum[i-1]+i)
				for j = i + i; j <= 1000000; j += i {
					np[j] = true
				}
			} else {
				sum = append(sum, sum[i-1])
			}
		}
		for t, _ := strconv.Atoi(scanner.Text()); t > 0 && scanner.Scan(); t-- {
			n, _ := strconv.Atoi(scanner.Text())
			wout.WriteString(strconv.FormatUint(sum[n], 10))
			wout.WriteString("\n")
		}
		wout.Flush()
	}
}
