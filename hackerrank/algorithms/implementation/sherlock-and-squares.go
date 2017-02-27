package main
import (
	"math"
	"fmt"
)
func main(){
	var t, a, b int
	fmt.Scanf("%d", &t)
	for ; t > 0; t-- {
		fmt.Scanf("%d %d", &a, &b)
		s := 0
		c := int(math.Sqrt(float64(a)))
		if a / c == c && a % c == 0 {
			s++
		}
		for i := c+1; i*i <= b; i++ {
			s++
		}
		fmt.Println(s)
	}
}
