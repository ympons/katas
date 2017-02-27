package main
import "fmt"

var t, n, m, s int
func main() {
    fmt.Scanf("%d", &t)
    for ;t>0;t--{
        fmt.Scanf("%d %d %d", &n, &m, &s)
        r := (s + m - 1)%n
        if r == 0 {
            r = n
        }
        fmt.Println(r)
    }
}
