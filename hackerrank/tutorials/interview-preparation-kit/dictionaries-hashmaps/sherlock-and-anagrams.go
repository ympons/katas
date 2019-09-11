package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func sherlockAndAnagrams(s string) int {
	n := len(s)
	pairs := 0
	for i := 1; i < n; i++ {
		hash := make(map[string]int)
		for j := 0; j <= n-i; j++ {
			substring := []rune(s[j : j+i])
			sort.Slice(substring, func(i, j int) bool {
				return substring[i] < substring[j]
			})
			if count, ok := hash[string(substring)]; ok {
				pairs += count
				hash[string(substring)]++
			} else {
				hash[string(substring)] = 1
			}
		}
	}
	return pairs
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var (
		q int
		s string
	)
	fmt.Fscanf(in, "%d\n", &q)
	for ; q > 0; q-- {
		fmt.Fscanf(in, "%s\n", &s)

		result := sherlockAndAnagrams(s)
		fmt.Fprintf(out, "%d\n", result)
	}
	out.Flush()
}
