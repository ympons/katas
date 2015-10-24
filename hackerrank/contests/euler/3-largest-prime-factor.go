package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func PrimeFactor(n float64) float64 {
	for i := 2.0; i <= math.Trunc(math.Sqrt(n)); i += 1.0 {
		if math.Mod(n, i) == 0.0 {
			return PrimeFactor(n / i)
		}
	}

	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wout := bufio.NewWriter(os.Stdout)
	if scanner.Scan() {
		for t, _ := strconv.Atoi(scanner.Text()); t > 0 && scanner.Scan(); t-- {
			n, _ := strconv.ParseFloat(scanner.Text(), 64)
			wout.WriteString(strconv.FormatFloat(PrimeFactor(n), 'f', -1, 64))
			wout.WriteString("\n")
		}
		wout.Flush()
	}

}
