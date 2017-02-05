package main
import "fmt"

var (
  t, n int
  a []int
)

func main(){
  fmt.Scanf("%d", &t)
  for ;t>0;t--{
    fmt.Scanf("%d", &n)
    a = make([]int, n)
    sum := 0
    for i := 0; i<n; i++{
      fmt.Scanf("%d", &a[i])
      sum += a[i]
    }
    if n == 1 {
      fmt.Println("YES")
    } else {
      ls, f := 0, false
      for i := 1; i<n; i++{
        ls += a[i-1]
        if (ls == sum-ls-a[i]) {
          f = true
          fmt.Println("YES")
          break
        }
      }
      if !f {
        fmt.Println("NO")
      }
    }
  }
}
