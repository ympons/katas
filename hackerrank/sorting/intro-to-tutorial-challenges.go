package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)
	if scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		if scanner.Scan() {
			n, _ := strconv.Atoi(scanner.Text())

			for i := 0; i < n && scanner.Scan(); i++ {
				a, _ := strconv.Atoi(scanner.Text())
				if a == v {
					out.WriteString(strconv.Itoa(i))
					break
				}
			}
		}
		out.Flush()
	}
}
