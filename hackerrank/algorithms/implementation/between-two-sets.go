package main
import (
    "fmt"
    "bufio"
    "os"
    "io"
)

func GCD(a, b int) int {
    if b == 0 {
        return a
    }
    return GCD(b, b%a)
}

func LCM(a, b int) int {
    return a*b/GCD(a,b)
}

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var n, m, x int
    fmt.Fscan(in, &n, &m)
    lcm := 1
    for i:=0;i<n;i++{
        fmt.Fscan(in, &x)
        lcm = LCM(x, lcm)
    }
    a := make([]int, m)
    for i:=0;i<m;i++{
        fmt.Fscan(in, &a[i])
    }
    
    sol := 0
    for i:=1;i<101;i++{
        auxS := 0
        for j:=0;j<m;j++{
            if (lcm * i > a[j]) {
                auxS = -1
                break
            }
            if a[j] % (lcm * i) != 0 {
                break
            }
            auxS++
        }
        if auxS == m {
            sol++
        } else if auxS == -1 {
            break
        }
    }
    fmt.Println(sol)
}
