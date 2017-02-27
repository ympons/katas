package main
import (
    "fmt"
    "os"
    "io"
    "bufio"
)

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var n, k, charge int
    fmt.Fscan(in, &n, &k)
    var sum, x int
    for i:=0;i<n;i++{
        fmt.Fscan(in, &x)
        if i != k {
            sum += x
        }
    }
    fmt.Fscan(in, &charge)
    if c := charge - sum/2; c <= 0 {
        fmt.Println("Bon Appetit")
    } else {
        fmt.Println(c)
    }
}
