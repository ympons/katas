package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	o := bufio.NewWriter(os.Stdout)
	r := int64(0)
	if s.Scan() {
		for n, _ := strconv.Atoi(s.Text()); n > 0 && s.Scan(); n-- {
			x, _ := strconv.ParseInt(s.Text(), 10, 64)
			r += x
		}
		o.WriteString(strconv.FormatInt(r, 10))
	}
	o.Flush()
}
