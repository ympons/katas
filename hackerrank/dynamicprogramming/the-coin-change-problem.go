package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)
	if scan.Scan() {
		n, _ := strconv.Atoi(scan.Text())
		if scan.Scan() {
			m, _ := strconv.Atoi(scan.Text())
			c, w := make([]int, 0, m), make([]uint64, n+1)
			for i := 0; i < m && scan.Scan(); i++ {
				a, _ := strconv.Atoi(scan.Text())
				c = append(c, a)
			}
			w[0] = 1
			for i := 0; i < m; i++ {
				for j := c[i]; j < n+1; j++ {
					w[j] += w[j - c[i]]
				}
			}
			out.WriteString(strconv.FormatUint(w[n],10))
		}
	}
	out.Flush()
}
