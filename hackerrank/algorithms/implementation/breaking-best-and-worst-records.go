package main
import ( 
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)
    
    var min int
    fmt.Scanf("%d", &min)
    
    max := min
    
    cMax, cMin := 0, 0
    for i := 1; i < n; i++{
        var x int
        fmt.Scanf("%d", &x)
        if x > max {
            max = x
            cMax++
        } else if x < min {
            min = x
            cMin++
        }
    }
    fmt.Println(cMax, cMin)
}
