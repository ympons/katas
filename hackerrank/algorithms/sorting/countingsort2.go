package main
import (
    "fmt"
    "io"
    "bufio"
    "os"
)

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var n, x int
    fmt.Fscan(in, &n)
    m := make([]int, 100)
    for i:=0;i<n;i++{
        fmt.Fscan(in, &x)
        m[x]++
    }
    for i:=0;i<100;i++{
        for j:=0;j<m[i];j++ {
            fmt.Printf("%d ", i)
        }
    }
}
