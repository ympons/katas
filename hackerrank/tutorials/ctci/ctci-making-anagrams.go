package main
import "fmt"

func main() {
    var a, b string
    fmt.Scan(&a)
    fmt.Scan(&b)
    m := make(map[rune]int)
    for _, v := range a {
        m[v]++
    }
    for _, v := range b {
        m[v]--
    }
    s := 0
    for _, v := range m {
        if v > 0 {
            s += v
        } else {
            s -= v
        }
    }
    fmt.Println(s)
}
