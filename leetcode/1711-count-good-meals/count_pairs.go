package leetcode

const module = int(1e9 + 7)

func countPairs(arr []int) int {
	M, n := map[int]int{}, len(arr)
	M[arr[0]]++

	answer := 0
	for i := 1; i < n; i++ {
		for j := 0; j < 31; j++ {
			answer = (answer + M[(1<<j)-arr[i]]) % module
		}
		M[arr[i]]++
	}

	return answer
}
