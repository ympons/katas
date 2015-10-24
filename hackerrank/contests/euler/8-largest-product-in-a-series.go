package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wout := bufio.NewWriter(os.Stdout)
	if scanner.Scan() {
		for t, _ := strconv.Atoi(scanner.Text()); t > 0 && scanner.Scan(); t-- {
			l := strings.Split(scanner.Text(), " ")
			n, _ := strconv.Atoi(l[0])
			k, _ := strconv.Atoi(l[1])
			var s uint64
			if scanner.Scan() {
				v := scanner.Text()
				for i := 0; i <= n-k; i++ {
					var p uint64 = 1
					for j := i; j <= i+k-1; j++ {
						p *= uint64(v[j] - '0')
					}
					if p > s {
						s = p
					}

				}
			}
			wout.WriteString(strconv.FormatUint(s, 10))
			wout.WriteString("\n")
		}
		wout.Flush()
	}
}
