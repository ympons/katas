package main
import "fmt"

type Node struct{
    value byte
    next *Node
}

func Add(nod *Node, v byte) *Node{
    return &Node{v, nod}
}

func Balanced(n *Node, s string) string {
    r := "NO"
    for i:=0;i<len(s);i++{
        switch s[i] {
            case '{','(','[':
                n = Add(n, s[i])
            case '}':
                if (n == nil || n.value != '{') {
                    return r
                } else {
                    n = n.next
                }
            case ')':       
                if (n == nil || n.value != '(') {
                    return r
                } else {
                    n = n.next
                }
            case ']':
                if (n == nil || n.value != '[') {
                    return r
                } else {
                    n = n.next
                }
        }
    }
    if (n != nil) {
        return r
    }
    return "YES"
}

func main() {
    var (
        t int
        s string
        n *Node
    )
    fmt.Scan(&t)
    for ;t>0;t--{
        fmt.Scan(&s)
        fmt.Println(Balanced(n, s))
    }
}
