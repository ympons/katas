package main
import "fmt"

func main() {
    var n, x, max, idx int
    m := make([]int, 6)
    fmt.Scan(&n)
    for i := 0; i<n; i++{
        fmt.Scan(&x)
        m[x]++
        if max < m[x] {
            max = m[x]
            idx = x
        } else if max == m[x] && x < idx {
            idx = x
        }
    }
    fmt.Println(idx)
}
