package main
import "fmt"

func main() {
    var n, k, x, max int
    fmt.Scanf("%d %d", &n, &k)
    for i:=0; i<n; i++{
        fmt.Scanf("%d", &x)
        if x > max {
            max = x
        }
    }
    x = 0
    if max > k {
        x = max - k
    }
    fmt.Println(x)
}
