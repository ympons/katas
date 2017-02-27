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
			s := in.Text()
			l, r := len(s), 0
			n, _ := strconv.Atoi(s)
			for i := 0; i < l; i++ {
				d := int(s[i] - 48)
				if (d > 0 && n % d == 0) {
					r++
				}
			}
			out.WriteString(strconv.Itoa(r))
			out.WriteString("\n")
		}
	}
	out.Flush()
}
