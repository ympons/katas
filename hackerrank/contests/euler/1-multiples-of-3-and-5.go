package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	// Just apply Sn = (n/2) * (a1 + an) and an = a1 + (n-1)d
	scanner := bufio.NewScanner(os.Stdin)
	wout := bufio.NewWriter(os.Stdout)
	if scanner.Scan() {
		var (
			n, s int
		)
		for t, _ := strconv.Atoi(scanner.Text()); t > 0 && scanner.Scan(); t-- {
			n, _ = strconv.Atoi(scanner.Text())
			n--
			// Sn3 + Sn5 - Sn15
			s3, s5, s15 := 3*(n/3)*(1+(n/3))/2, 5*(n/5)*(1+(n/5))/2, 15*(n/15)*(1+(n/15))/2
			s = s3 + s5 - s15
			wout.WriteString(strconv.Itoa(s))
			wout.WriteString("\n")
		}
		wout.Flush()
	}

}
