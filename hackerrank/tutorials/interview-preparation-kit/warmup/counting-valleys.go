package main

import "fmt"

func countingValleys(n int32, s string) int32 {
	steps, valleys := int32(0), int32(0)
	for _, c := range s {
		if c == 'U' {
			if steps == -1 {
				valleys++
			}
			steps++
		} else {
			steps--
		}
	}
	return valleys
}

func main() {
	var (
		n int32
		s string
	)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &s)

	result := countingValleys(n, s)
	fmt.Printf("%d", result)
}
