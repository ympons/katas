package leetcode

const module = int(1e9 + 7)

func concatenatedBinary(n int) int {
	answer, size := 0, 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			size++
		}
		answer = (answer*(1<<size) + i) % module
	}
	return answer
}
