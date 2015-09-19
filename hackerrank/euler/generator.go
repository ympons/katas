package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var l []int
var m map[int]bool

func Palin(n int) bool {
	x := strconv.Itoa(n)
	s := len(x)
	for i := 0; i < s/2; i++ {
		if x[i] != x[s-i-1] {
			return false
		}
	}
	return true
}

func main() {
	f, _ := os.Create("test.txt")
	out := bufio.NewWriter(f)
	defer f.Close()

	l = make([]int, 0, 100000)
	m = make(map[int]bool)

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			x := i * j
			if x >= 101101 && Palin(x) {
				if !m[x] {
					l = append(l, x)
					m[x] = true
				}
			}
		}
	}

	sort.Ints(l)
	for i := 0; i < len(l); i++ {
		out.WriteString(strconv.Itoa(l[i]))
		out.WriteString(",")
	}
	out.Flush()
}
