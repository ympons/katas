package main

import (
	"bufio"
	"os"
	"strconv"
)

func GCD(a, b uint64) uint64 {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

func LCM(a, b uint64) uint64 {
	return a * b / GCD(a, b)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wout := bufio.NewWriter(os.Stdout)
	if scanner.Scan() {
		for t, _ := strconv.Atoi(scanner.Text()); t > 0 && scanner.Scan(); t-- {
			n, _ := strconv.Atoi(scanner.Text())
			var s uint64 = 1
			for i := 1; i <= n; i++ {
				s = LCM(uint64(i), s)
			}
			wout.WriteString(strconv.FormatUint(s, 10))
			wout.WriteString("\n")
		}
		wout.Flush()
	}
}
