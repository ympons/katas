package main
import (
    "fmt"
    "bufio"
    "io"
    "os"
    "sort"
)

func insertionSort(a []int) int {
    swaps := 0;
    for i := 1; i < len(a); i++ {
        for j := i; j > 0 && a[j] < a[j-1]; j-- {
            a[j], a[j-1] = a[j-1], a[j]
            swaps++
        }
    }
    return swaps
}

func quickSort(a sort.Interface, left, right int) int {
	if left >= right {
		return 0
	}
	piv, swaps := partition(a, left, right, right)
	return swaps + quickSort(a, left, piv-1) + quickSort(a, piv+1, right)

}

func partition(a sort.Interface, left int, right int, piv int) (int, int) {
	j, swaps := left, 0
	for i := left; i <= right; i++ {
		if a.Less(i, piv) {
			a.Swap(i, j)
			j++
			swaps++
		}
	}
	a.Swap(piv, j)
	swaps++
	return j, swaps
}

var in io.Reader = bufio.NewReader(os.Stdin)
func main() {
    var n int
    fmt.Fscan(in, &n)
    a, b := make([]int, n), make([]int, n)
    for i:=0;i<n;i++{
        fmt.Fscan(in, &a[i])
    }
    copy(b, a)
    fmt.Println(insertionSort(a) - quickSort(sort.IntSlice(b), 0, n-1))
}
