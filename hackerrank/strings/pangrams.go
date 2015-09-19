package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	o := bufio.NewWriter(os.Stdout)
	i := bufio.NewScanner(os.Stdin)
	i.Scan()
	m := make(map[interface{}]bool)
	s := strings.Replace(strings.ToLower(i.Text()), " ", "", -1)
	for _, x := range s {
		m[x] = true
	}
	if len(m) == 26 {
		o.WriteString("pangram")
	} else {
		o.WriteString("no pangram")
	}
	o.Flush()
}
