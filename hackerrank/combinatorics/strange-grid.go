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
		r, _ := strconv.ParseUint(scan.Text(),10,64)
		if scan.Scan() {
			c, _ := strconv.ParseUint(scan.Text(),10,64)
			if r % 2 == 0 {
				out.WriteString(strconv.FormatUint((r-2)*5 + 2*(c-1) + 1, 10))
			} else {
				out.WriteString(strconv.FormatUint((r-1)*5 + 2*(c-1), 10))
			}
		}
	}
	out.Flush()
}
