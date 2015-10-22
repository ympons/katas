package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func N(s string) uint64 {
	x, _ := strconv.Atoi(s)
	return uint64(x)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	g := make([][]string, 20)
	for i := 0; scanner.Scan() && i < 20; i++ {
		l := strings.Split(scanner.Text(), " ")
		g[i] = append(g[i], l...)
	}
	var p, s uint64
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			if j+3 < 20 {
				p = N(g[i][j]) * N(g[i][j+1]) * N(g[i][j+2]) * N(g[i][j+3])
				if p > s {
					s = p
				}
			}
			if i+3 < 20 {
				p = N(g[i][j]) * N(g[i+1][j]) * N(g[i+2][j]) * N(g[i+3][j])
				if p > s {
					s = p
				}
			}

			if i+3 < 20 && j+3 < 20 {
				p = N(g[i][j]) * N(g[i+1][j+1]) * N(g[i+2][j+2]) * N(g[i+3][j+3])
				if p > s {
					s = p
				}
			}
			if i+3 < 20 && j-3 >= 0 {
				p = N(g[i][j]) * N(g[i+1][j-1]) * N(g[i+2][j-2]) * N(g[i+3][j-3])
				if p > s {
					s = p
				}
			}
		}
	}
	out.WriteString(strconv.FormatUint(s, 10))
	out.Flush()
}
