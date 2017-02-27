package main

import (
	"fmt"
	"math"
)

func main() {
	var t float64
	fmt.Scan(&t)
	var c int64 = int64(3*(math.Pow(2, math.Ceil(math.Log2(t/3 + 1)))-1) + 1 - t)
	fmt.Print(c)
}

