package main
import "fmt"

func main() {
    var (
        x1,x2,v1,v2 int
    )
    fmt.Scanf("%d%d%d%d",&x1,&v1,&x2,&v2)
    x := x2-x1
    v := v1-v2
    if v != 0 && x % v == 0 && x/v >= 0 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
