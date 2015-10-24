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
		t, _ := strconv.Atoi(scan.Text())
		for i := 0; i < t && scan.Scan(); t-- {
			n, _ := strconv.ParseUint(scan.Text(), 10, 64)
			out.WriteString(strconv.FormatUint(n*(n-1)/2, 10))
			out.WriteString("\n")
		}
	}
	out.Flush()
}
