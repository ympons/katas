package main
import "fmt"

var n, k, s int
func main() {
    fmt.Scanf("%d %d", &n, &k)
    L := make([]int, n)
    for i:=0;i<n;i++{
        fmt.Scanf("%d", &L[i])
    }
    s = 0
    for i:=0;i<n-1;i++{
        for j:=i+1;j<n;j++{
            if (L[i]+L[j])%k==0 {
                s++
            }
        }
    }
    fmt.Println(s)
}
