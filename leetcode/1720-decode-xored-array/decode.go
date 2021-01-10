package leetcode

func decode(encoded []int, first int) []int {
	n := len(encoded) + 1
	A := make([]int, n)
	A[0] = first
	for i := 1; i < n; i++ {
		A[i] = A[i-1] ^ encoded[i-1]
	}
	return A
}
