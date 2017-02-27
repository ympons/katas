package main

import (
	"fmt"
	"sort"
)

var (
	n int
	a sort.IntSlice
)

func main(){
	fmt.Scanf("%d", &n)
	a = make(sort.IntSlice, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	sort.Sort(a)
	fmt.Println(n)
	for i, acc, r := 0, 0, n; i < n; {
		r -= 1
		j := i+1
		for ;j < n && ((a[i]-acc) == (a[j]-acc)); j++ { r-- }
		acc = a[i]
		i = j
		if r > 0 { fmt.Println(r) }
	}
}
