package main

import (
	"bufio"
	"fmt"
	"os"
)

func repeatedString(s string, n int64) int64 {
	count := int64(0)
	for _, c := range s {
		if c == 'a' {
			count++
		}
	}
	size := int64(len(s))
	count = count * (n / size)
	rem := int(n % size)
	for i := 0; i < rem; i++ {
		if s[i] == 'a' {
			count++
		}
	}
	return count
}

func main() {
	var (
		n int64
		s string
	)
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &s)
	fmt.Fscan(in, &n)

	result := repeatedString(s, n)
	fmt.Println(result)
}
