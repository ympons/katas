package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func minimumBribes(arr []int) {
	for i, x := range arr {
		if x-i-1 > 2 {
			out.WriteString("Too chaotic\n")
			return
		}
	}

	bribes := 0
	for sorted := false; !sorted; {
		sorted = true
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				sorted = false
				bribes++
			}
		}
	}
	out.WriteString(strconv.FormatUint(uint64(bribes), 10))
	out.WriteString("\n")
}

func main() {
	var t, n int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		fmt.Fscan(in, &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		minimumBribes(arr)
	}
	out.Flush()
}
