package main
import (
	"fmt"
    "os"
    "io"
    "bufio"
)

var (
	t int
	w string
)

func next(s []uint8) []uint8 {
    // Find non-increasing suffix
    i := len(s) - 1;
    for ;i>0 && s[i-1] >= s[i]; i--{}
    if (i <= 0) {
        return []uint8("no answer")
    }
    // Find successor to pivot
    j := len(s) - 1;
    for ;s[j] <= s[i - 1]; j--{}

    tt := s[i - 1]
    s[i - 1], s[j] = s[j], tt

    // Reverse suffix
    j = len(s) - 1;
    for i < j {
        tt = s[i]
        s[i], s[j] = s[j], tt
        i++
        j--
    }
    return s
}

var in io.Reader = bufio.NewReader(os.Stdin)
func main(){
	fmt.Fscan(in, &t)
	for ;t>0;t--{
		fmt.Fscan(in, &w)
		if len(w) == 1 {
			fmt.Println("no answer")
		} else {
			fmt.Printf("%s\n", next([]uint8(w)))
		}
	}
}
