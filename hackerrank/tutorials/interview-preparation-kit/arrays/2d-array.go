package main

import "fmt"

func main() {
	a := make([][]int, 6)
	for i := 0; i < 6; i++ {
		a[i] = make([]int, 6)
		for j := 0; j < 6; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}
	sol := -360
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s := (a[i][j] + a[i][j+1] + a[i][j+2])
			s += (a[i+1][j+1])
			s += (a[i+2][j] + a[i+2][j+1] + a[i+2][j+2])
			if s > sol {
				sol = s
			}
		}
	}
	fmt.Printf("%d", sol)
}
