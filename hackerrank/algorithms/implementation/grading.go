package main
import (
    "fmt"
    "io"
    "bufio"
    "os"
)

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var t, x int
    fmt.Fscan(in, &t)
    for ;t>0;t--{
        fmt.Fscan(in, &x)
        if x < 38 || x % 5 == 0 {
            fmt.Println(x)
        } else {
            next := (x/5 + 1)*5
            if next - x < 3 {
                fmt.Println(next)
            } else {
                fmt.Println(x)
            }
        }
    }
}
