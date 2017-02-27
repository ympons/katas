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
		t, _ := strconv.Atoi(in.Text())
		p, n, z := 0, 0, 0
		for i := 0; i < t && in.Scan(); i++ {
			x, _ := strconv.Atoi(in.Text())
			if x < 0 {
				n++
			} else if x == 0 {
				z++
			} else {
				p++
			}
		}
		out.WriteString(strconv.FormatFloat(float64(p)/float64(t), 'f', 6, 64))
		out.WriteString("\n")
		out.WriteString(strconv.FormatFloat(float64(n)/float64(t), 'f', 6, 64))
		out.WriteString("\n")
		out.WriteString(strconv.FormatFloat(float64(z)/float64(t), 'f', 6, 64))
	}
	out.Flush()
}
