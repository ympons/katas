package main
import "fmt"

func main() {
    var n,d,b int
    fmt.Scan(&n, &d)
    a := make([]int, n)
    m := make(map[int]bool)
    for i:=0;i<n;i++{
        fmt.Scan(&a[i])
        m[a[i]] = true
    }
    for i:=0;i<n-2;i++{
        if m[a[i]+d] && m[a[i]+2*d] {
            b++
        }
    }
    fmt.Println(b)
}
