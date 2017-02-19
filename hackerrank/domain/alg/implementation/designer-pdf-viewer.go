package main
import (
    "bufio"
    "fmt"
    "os"
    "io"
)

var in io.Reader = bufio.NewReader(os.Stdin)

func main() {
    h := make(map[rune]int)
    var x int
    for i:='a';i<='z';i++{
        fmt.Fscan(in, &x)
        h[i] = x
    }
    var s string
    hw := 1
    fmt.Fscan(in, &s)
    for _, x := range s {
        if h[x] > hw {
            hw = h[x]
        }
    }
    fmt.Println(hw * len(s))
}
