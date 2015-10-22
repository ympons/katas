package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	if scanner.Scan() {
		for t, _ := strconv.Atoi(scanner.Text()); t > 0 && scanner.Scan(); t-- {
			n, _ := strconv.Atoi(scanner.Text())
			var (
				a, b, c int
				s       int64 = -1
			)
			for b = 1; b <= n/2; b++ {
				if (n*n/2-b*n)%(n-b) == 0 {
					a = (n*n/2 - b*n) / (n - b)
					if a <= 0 || c > n {
						continue
					}
					c = int(math.Sqrt(float64(a*a + b*b)))
					if a+b+c == n && int64(a*b*c) > s {
						s = int64(a * b * c)
					}
				}
			}
			out.WriteString(strconv.FormatInt(s, 10))
			out.WriteString("\n")
		}
		out.Flush()
	}
}
