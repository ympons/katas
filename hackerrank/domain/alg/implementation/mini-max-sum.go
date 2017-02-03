package main
import "fmt"

func main() {
    var x int
    max, min, sum := 0, 1000000005, 0
    for i:=0;i<5;i++{
        fmt.Scan(&x)
        sum += x
        if x < min {
            min = x
        } 
        if x > max {
            max = x
        }
    }
    fmt.Print(sum-max, sum-min)
}
