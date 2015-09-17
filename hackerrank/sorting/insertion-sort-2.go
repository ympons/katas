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
		l := make([]int, 0, n)
		var (
			i, a int
			p    = func() {
				for i := 0; i < n; i++ {
					out.WriteString(strconv.Itoa(l[i]))
					out.WriteString(" ")
				}
				out.WriteString("\n")
			}
		)
		for i = 0; i < n && scan.Scan(); i++ {
			a, _ := strconv.Atoi(scan.Text())
			l = append(l, a)
		}
		for k := 1; k < n; k++ {
			a = l[k]
			for i = k - 1; i >= 0 && l[i] > a; i-- {
				l[i+1] = l[i]
			}
			l[i+1] = a
			p()
		}
	}
	out.Flush()
}
