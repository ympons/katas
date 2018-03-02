package hamming

import "errors"

var ErrStrandsInvalidSize = errors.New("hamming: DNA strands must have equal size")

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, ErrStrandsInvalidSize
	}
	distance := 0
	for i := range a {
		if b[i] != a[i] {
			distance++
		}
	}
	return distance, nil
}
