package main
import (
    "bufio"
    "io"
    "os"
    "fmt"
)

var in io.Reader = bufio.NewReader(os.Stdin)

func main() {
    var n, x int
    fmt.Fscan(in, &n)
    l := make([]int, 10001)
    for i:=0;i<n;i++{
        fmt.Fscan(in, &x)
        l[x]++
    }
    fmt.Fscan(in, &n)
    for i:=0;i<n;i++{
        fmt.Fscan(in, &x)
        l[x]--
    }
    for i:=1; i<10001; i++{
        if l[i]!=0 {
            fmt.Printf("%d ", i)
        }
    }
}
