package main
import "fmt"

var (
    n int
    a []int
)
func main() {
    fmt.Scanf("%d", &n)
    a = make([]int, n)
    for i:=0;i<n;i++{
        fmt.Scanf("%d", &a[i])
    }
    s := 0
    for i:=0;i<n;{
        if (i+2 < n && a[i+2] != 1) {
           i += 2
        } else {
           i++
        }
        s++
    }
    fmt.Printf("%d", s-1)
}
