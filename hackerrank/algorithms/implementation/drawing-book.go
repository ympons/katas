package main

import "fmt"

func main() {
	var n, p int
	fmt.Scanf("%d\n%d", &n, &p)
	cn, cp := n/2, p/2
	if x := cn - cp; x < cp {
		cp = x
	}
	fmt.Println(cp)
}
