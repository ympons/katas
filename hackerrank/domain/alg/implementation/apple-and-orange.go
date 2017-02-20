package main
import (
    "fmt"
    "os"
    "bufio"
    "io"
)

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var (
        s, t, a, b, m, n int
    )
    fmt.Fscan(in, &s, &t)
    fmt.Fscan(in, &a, &b)
    fmt.Fscan(in, &m, &n)
    apples, oranges := 0, 0
    var x int
    for i:=0;i<m;i++{
        fmt.Fscan(in, &x)
        if v := a + x; v >= s && v <= t {
            apples++
        }
    }
    for i:=0;i<n;i++{
        fmt.Fscan(in, &x)
        if v := b + x; v >= s && v <= t {
            oranges++
        }
    }
    fmt.Println(apples)
    fmt.Println(oranges)    
}
