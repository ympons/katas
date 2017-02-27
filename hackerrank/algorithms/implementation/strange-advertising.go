package main
import (
    "fmt"
    "os"
    "io"
    "bufio"
)

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var n int
    fmt.Fscan(in, &n)
    sum, p := 0, 5
    for i:=0;i<n;i++{
        p /= 2
        sum += p
        p *= 3
    }
    fmt.Println(sum)
}
