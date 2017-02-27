package main
import "fmt"

var (
  t, m, n int
  l []int
)

func main(){
  fmt.Scanf("%d", &t)
  solve := func() (int,int) {
    for i:=0;i<n-1;i++{
      for j:=i+1;j<n;j++{
        if l[i] + l[j] == m {
          return i+1, j+1
        }
      }
    }
    return 0, 0
  }
  for ;t>0;t--{
    fmt.Scanf("%d", &m)
    fmt.Scanf("%d", &n)
    l = make([]int, n)
    for i:=0;i<n;i++{
      fmt.Scanf("%d", &l[i])
    }
    i, j := solve()
    fmt.Printf("%d %d\n",i,j)
  }
}
