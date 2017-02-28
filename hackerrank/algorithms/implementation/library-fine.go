package main
import (
    "fmt"
    "bufio"
    "io"
    "os"
)

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var (
        da,ma,ya int
        db,mb,yb int
    )
    fmt.Fscan(in, &db, &mb, &yb)
    fmt.Fscan(in, &da, &ma, &ya)
    fine := 0
    if yb > ya {
        fine = 10000
    } else if yb == ya {
        if mb > ma {
            fine = 500*(mb-ma)
        } else if mb == ma {
            if db > da {
                fine = 15*(db-da)
            }
        }
    }
    fmt.Println(fine)
}
