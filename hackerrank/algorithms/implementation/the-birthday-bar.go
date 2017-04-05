package main
import (
    "fmt"
    "io"
    "os"
    "bufio"
)

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
	var n, d, m int

	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Fscan(in, &d, &m)

	sum, count := 0, 0
	if m <= n {
		for i := 0; i < m; i++ {
			sum += a[i]
		}
		if sum == d {
			count++
		}

		for prev, i, j := 0, 1, m; j < n; i, j = i+1, j+1 {
			sum = sum - a[prev] + a[j]
			if sum == d {
				count++
			}
			prev = i
		}
	}
	fmt.Println(count)
}
