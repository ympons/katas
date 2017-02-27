package main
import (
    "fmt"
    "io"
    "bufio"
    "os"
)

func reverse(x int) int {
    r := 0
    for x != 0 {
        r =  (r + x%10) * 10
        x /= 10        
    }
    return r/10
}
func abs(x int) int {
    if x < 0 {
        return -1 * x
    }
    return x
}

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var i, j, k int
    fmt.Fscan(in, &i, &j, &k)
    sol := 0
    for ;i <= j; i++ {
        if abs(i - reverse(i)) % k == 0 {
            sol++
        }
    }
    fmt.Println(sol)
}
