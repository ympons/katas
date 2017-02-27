package main
import (
    "fmt"
    "bufio"
    "io"
    "os"
)

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var max, x, n int
    fmt.Fscan(in, &n)
    m := make(map[int]int)
    for i:=0;i<n;i++{
        fmt.Fscan(in, &x)
        m[x]++
        if max < m[x] {
            max = m[x]
        }
    }
    fmt.Println(n-max)
}
