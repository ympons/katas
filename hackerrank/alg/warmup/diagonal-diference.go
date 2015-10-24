package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	o := bufio.NewWriter(os.Stdout)
	i := bufio.NewScanner(os.Stdin)
	i.Scan()
	n, _ := strconv.Atoi(i.Text())
	r := 0
	for j := 0; j < n && i.Scan(); j++ {
		s := strings.Split(i.Text(), " ")
		a, _ := strconv.Atoi(s[j])
		b, _ := strconv.Atoi(s[n-j-1])
		r += (a - b)
	}
	if r < 0 {
		r *= -1
	}
	o.WriteString(strconv.Itoa(r))
	o.Flush()
}
