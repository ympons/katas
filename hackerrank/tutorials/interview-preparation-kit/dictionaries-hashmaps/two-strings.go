package main

import (
	"bufio"
	"fmt"
	"os"
)

func twoStrings(s1 string, s2 string) string {
	hash := make(map[rune]bool)
	for _, c := range s1 {
		hash[c] = true
	}
	for _, c := range s2 {
		if hash[c] {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var (
		p      int
		s1, s2 string
	)
	fmt.Fscanf(in, "%d\n", &p)
	for i := 0; i < p; i++ {
		fmt.Fscanf(in, "%s\n", &s1)
		fmt.Fscanf(in, "%s\n", &s2)

		result := twoStrings(s1, s2)
		fmt.Fprintf(out, "%s\n", result)
	}
	out.Flush()
}
