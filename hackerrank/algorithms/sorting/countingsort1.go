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
    m := make(map[int]int)
    for i:=0;i<n;i++{
        fmt.Fscan(in, &x)
        m[x]++
    }
    for i:=0;i<100;i++{
        fmt.Printf("%d ", m[i])
    }
}
