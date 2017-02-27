package main
import "fmt"

var (
    n,k int
    a []int
)
func main() {
    fmt.Scanf("%d%d", &n,&k)
    a = make([]int, n)
    for i:=0;i<n;i++{
        fmt.Scanf("%d", &a[i])
    }
    e := 100
    for i:=0;i==0||i%n!=0;i+=k{
        if a[i] == 1 {
            e-=3
        } else {
            e--
        }
    }
    fmt.Printf("%d", e)
}
