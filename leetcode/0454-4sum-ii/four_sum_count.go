package leetcode

func fourSumCount(A []int, B []int, C []int, D []int) int {
	counter := map[int]int{}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			counter[A[i]+B[j]]++
		}
	}
	answer := 0
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			sum1, sum2 := C[i]+D[j], 0
			if sum1 != 0 {
				sum2 = -sum1
			}
			if c, ok := counter[sum2]; ok {
				answer += c
			}
		}
	}
	return answer
}
