package main
import (
	"fmt"
	"sort"
)

type (
	P struct {
		idx, v int
	}
	L []P
)
func (c L) Len() int { return len(c) }
func (c L) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c L) Less(i, j int) bool { return c[i].v < c[j].v || (c[i].v == c[j].v && c[i].idx < c[j].idx) }

var n int

func main(){
	fmt.Scanf("%d", &n)
    l := make(L, n)
    for i:=0;i<n;i++{
        fmt.Scanf("%d", &l[i].v)
        l[i].idx = i
    }
    sort.Sort(l)
    sol := n+1
    for i := 0;i<n;{
        j := i+1
        for ;j < n && l[i].v==l[j].v;j++{}
        if j > i+1 && j <= n {
            if s := l[i+1].idx - l[i].idx; sol > s {
                sol = s
            }
        }
        i = j
    }
    if sol == n+1 {
        fmt.Print("-1")
    } else {
        fmt.Print(sol)
    }
}
