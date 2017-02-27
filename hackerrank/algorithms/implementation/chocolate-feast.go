package main
import "fmt"

var t,n,c,m int

func main(){
	fmt.Scanf("%d", &t)
	for ;t>0;t--{
		fmt.Scanf("%d %d %d", &n, &c, &m)
		r := n/c
		total := r
		for r >= m {
			n = r/m
			total += n
			r = n + r%m
		}
		fmt.Println(total)
	}
}
