package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	if scanner.Scan() {
		for t, _ := strconv.Atoi(scanner.Text()); t > 0 && scanner.Scan(); t-- {
			l := strings.Split(scanner.Text(), " ")
			a, _ := strconv.Atoi(l[0])
			b, _ := strconv.Atoi(l[1])
			var r float64
			for i := a; i <= b; i++ {
				r += math.Pow(float64(10), float64(i))
			}
			if r > 1000000 {
				out.WriteString("YES\n")
			} else {
				out.WriteString("NO\n")
			}
		}
		out.Flush()
	}
}
