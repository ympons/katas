package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)
	if in.Scan() {
		abs := func(v int) int {
			if v < 0 {
				return -v
			}
			return v
		}
		funny := func(s string) bool {
			n := len(s)
			for j := 1; j < n; j++ {
				if (abs(int(s[j])-int(s[j-1])) != abs(int(s[n-j])-int(s[n-j-1]))) {
					return false
				}
			}
			return true
		}
		for t, _ := strconv.Atoi(in.Text()); t > 0 && in.Scan(); t-- {
			if funny(in.Text()) {
				out.WriteString("Funny")
			} else {
				out.WriteString("Not Funny")
			}
			out.WriteString("\n")
		}
	}
	out.Flush()
}
