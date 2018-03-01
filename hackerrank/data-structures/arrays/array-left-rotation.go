package main
import (
    "bufio"
    "fmt"
    "os"
    "io"
)

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var n, d int
    fmt.Fscan(in, &n, &d)
    a := make([]int, n)
    for i:=0;i<n;i++{
        fmt.Fscan(in, &a[i])
    }
    d = d % n
    for i:=0;i<n;i++{
        fmt.Printf("%d ", a[(i + d)%n])
    }
}
