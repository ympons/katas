package main
import (
    "fmt"
    "os"
    "io"
    "bufio"
)

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var (
        s, t string
        k int
    )
    fmt.Fscan(in, &s)
    fmt.Fscan(in, &t)
    fmt.Fscan(in, &k)
    sn, st := len(s), len(t)
    i := 0
    for ;i<sn && i<st && s[i]==t[i]; i++{}
    if d := sn+st-2*i; k >= sn + st || ( k >= d && (k-d)%2 == 0) {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
